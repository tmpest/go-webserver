# go-webserver

This is my Golang Webserver for my Personal Website. I choose Golang to learn Go as it's something my Company is making the switch to and I wanted to get a head start on understanding it.

For now I have also choosen to try and deploy it using Kubernetes. This is in all reality overkill but again it was done to practice a technology my Company uses and get more familiar with it. In the future (when my Google Cloud free trial ends) I plan to take my container else where and deploy it using other means.

[The actual site](https://tmpest.com/)

## Relasing a new version of the Website

I plan to fully auto mate this but at a high level the steps are:
1. Build and tag a docker image
2. Publish the docker image to dockerhub
3. Publish the same image to dockerhub as the `:latest`
4. Connect to kubernetes cluster and edit the deployment `deployment-label` to anything other than what it currently is

Steps 1-4 happen automatically when code is pushed to the repository with a version.

Kicking off a kubernetes deployment can be done a few ways.
1. Delete the pod
    * Because the Deployment is configured to always pull the latest image deleting the pod will make kubernetes spin up a new pod, which will pull the latest code.
    * Not necessarily ideal since it will temporarily cause an outtage

2. Slightly better, update the deployment manifest, which triggers a Rolling Update
    * Updating the deployment manifest causes Kubernetes to try to correct the state of things. Updating the `deployment-label` effectively does nothing, it's a label I created just to be able to edit something that has no dependencies.
    * The result is kubernetes should spin up a new pod before deleting the old pod, which is much better

## V2 Infrastructure
I have move to a much cheaper setup and manually deploy my site on a server hosted by DigitalOcean

### Webserver
Starting the webserver can be done by running:
```
systemctl start webserver
```
It should start at boot though, in case it's not:
```
systemctl enable webserver
```
The `webserver.service` file is stored at `/etc/systemd/system/webserver.service`

### SSL
To get SSL this time I am using `nginx` and configuring a reverse proxy to the webserver which is running on port `8080`

The nginx server does a few things, it allows me to expose only ports `:80` and `443` for http and https access without exposing my webserver directly. It also will handle certificate verification which is the key component of tls.

the nginx server is started by running:
```
systemctl start nginx
```
but should boot at startup, if it's not run:
```
systemctl enable nginx
```

the configuration is located in `/etc/nginx/sites-available/joshnolivia.com`

changes to the configuration can be reloaded by running:
```
systemctl reload nginx
```

Certificates expire and need to be reloaded. A tool called `certbot` manages this process by rotating the certificates for domains defined.

[Digital Ocean's Documentation on how to do this](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04)

