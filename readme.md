# README

## Overview
This application scrapes the "Mediawatch" section of ABC News and generates an RSS feed. It supports both CLI and CGI modes.

## Installation

### Prerequisites
- **Go** (if building from source)

#### Optional
- **Apache or Nginx** (for serving RSS files)

### Build and Install

#### Install only (System level)
Grab the latest binary here: https://github.com/arran4/abc-mediawatch-rss/releases/

#### Install and build as user (User)
Install go 1.23+

Run `go install`:
```bash
go install github.com/arran4/abc-mediawatch-rss/cmd/abcmediawatchrss@latest
```
This installs to `$HOME/go/bin` (typically; check with `go env`).

### Usage
#### CLI Mode
Generate RSS Feed:
```bash
abcmediawatchrss -output /var/www/localhost/htdocs/rss/abcmediawatchrss.xml
```

#### CGI Mode
1. Place `abcmediawatchrss-cgi` in your server's CGI directory (e.g., `/var/www/htdocs/cgi-bin/abcmediawatchrss-cgi`).
2. Ensure it is executable:
   ```bash
   chmod +x /var/www/htdocs/cgi-bin/abcmediawatchrss-cgi
   ```
3. Access it via URL (e.g., `http://example.com/cgi-bin/abcmediawatchrss-cgi`).

### Deployment

#### rc.d (Cron Job system level)
Add a cron job to run the script periodically:
1. Edit the root crontab:
   ```bash
   sudo crontab -e
   ```
2. Add the following line:
   ```bash
   */15 * * * * /usr/local/bin/abcmediawatchrss -output /var/www/localhost/htdocs/rss/abcmediawatchrss.xml
   ```

#### rc.d (Cron Job user level)
Add a cron job to run the script periodically:
1. Edit the user's crontab:
   ```bash
   crontab -e
   ```
2. Add the following line:
   ```bash
   */15 * * * * ~/go/bin/abcmediawatchrss -output ~/public_html/rss/abcmediawatchrss.xml
   ```

#### systemd (as root)
1. Create a systemd service file at `/etc/systemd/system/abcmediawatchrss.service`:
```ini
[Unit]
Description=ABC Mediawatch RSS Feed Creator

[Service]
Type=oneshot
ExecStart=/usr/bin/abcmediawatchrss -output /var/www/localhost/htdocs/rss/abcmediawatchrss.xml
User=apache
Group=apache
```

2. Create a systemd timer file at `/etc/systemd/system/everyhour@.timer`:

```ini
[Unit]
Description=Monthly Timer for %i service

[Timer]
OnCalendar=*-*-* *:00:00
AccuracySec=1h
RandomizedDelaySec=1h
Persistent=true
Unit=%i.service

[Install]
WantedBy=default.target
```

3. Reload systemd and start the service:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable --now everyhour@abcmediawatchrss.timer
   ```

#### systemd (as user)
1. Create a systemd service file at `$HOME/.config/systemd/user/abcmediawatchrss.service`:
```ini
[Unit]
Description=ABC Mediawatch RSS Feed Creator

[Service]
Type=oneshot
ExecStart=%h/go/bin/abcmediawatchrss -output %h/public_html/rss/abcmediawatchrss.xml
```

2. Create a systemd timer file at `$HOME/.config/systemd/user/everyhour@.timer`:

```ini
[Unit]
Description=Monthly Timer for %i service

[Timer]
OnCalendar=*-*-* *:00:00
AccuracySec=1h
RandomizedDelaySec=1h
Persistent=true
Unit=%i.service

[Install]
WantedBy=default.target
```

3. Reload systemd and start the service:
   ```bash
   systemctl --user daemon-reload && systemctl --user enable --now everyhour@abcmediawatchrss.timer
   ```

#### Apache VirtualHost Configuration
##### User

Refer to documentation for setting up public_html directories

##### Enjoy

http://localhost/~$USERNAME/rss/abcmediawatchrss.xml

##### System

Add the following configuration to your Apache setup (e.g., `/etc/httpd/conf.d/rss.conf`):
```apache
<VirtualHost *:80>
    ServerName example.com
    DocumentRoot /var/www/localhost/htdocs/rss
    <Directory "/var/www/localhost/htdocs/rss">
        Options Indexes FollowSymLinks
        AllowOverride None
        Require all granted
    </Directory>
</VirtualHost>
```

#### Nginx Configuration
##### User

Refer to documentation for setting up public_html directories

##### System

Add this to your Nginx server block:
```nginx
server {
    listen 80;
    server_name example.com;

    location /rss/ {
        root /var/www/localhost/htdocs;
        autoindex on;
    }
}
```
