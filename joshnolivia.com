server {
        listen 80;
        listen [::]:80;

        root /var/www/joshnolivia.com/html;

        server_name joshnolivia.com www.joshnolivia.com;

        location / {
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $remote_addr;
                proxy_set_header Host      $host;
                proxy_pass http://localhost:8080;
        }
} #certbot nginx -d joshnolivia.com -d www.joshnolivia.com