## 1. Set up the Ansible project:
- Create a new directory for your Ansible project: **`ansible-webserver`.**
- Change to the **`ansible-webserver directory`.**
- Initialize the project with necessary files using the command: **`ansible-galaxy init roles/webserver.`**

Run the following commands in your terminal:
<details>
  <summary>bash</summary>

```json
$ mkdir ansible-webserver
$ cd ansible-webserver
$ ansible-galaxy init roles/webserver
```
</details>

## 2. Create the Ansible playbook:

- Inside the **`ansible-webserver`** directory, create a file called site.yml.
- Add the following content to **`site.yml`**:

<details>
  <summary>site.yml</summary>

```json
---
- name: Configure and deploy web server
  hosts: your_server
  become: yes
  roles:
    - webserver
```
</details>

## 3. Create the Ansible role:

- Navigate to the **`roles`** directory.
- Go to **`webserver`** and modify the **`tasks/main.yml`** file as follows:

<details>
  <summary>tasks/main.yml</summary>

```json
---
- name: Install required packages
  apt:
    name: "{{ item }}"
    state: present
  loop:
    - apache2
    - libapache2-mod-proxy-html
    - libxml2-dev
    - openssl

- name: Enable required Apache modules
  apache2_module:
    name: "{{ item }}"
    state: present
  loop:
    - proxy
    - proxy_http
    - proxy_html
    - ssl
    - rewrite

- name: Create self-signed SSL certificate
  openssl_certificate:
    path: /etc/ssl/certs/selfsigned.crt
    privatekey_path: /etc/ssl/private/selfsigned.key
    common_name: your_domain

- name: Configure Apache virtual host
  template:
    src: vhost.conf.j2
    dest: /etc/apache2/sites-available/your_domain.conf
    owner: root
    group: root
    mode: 0644
  notify:
    - restart Apache

- name: Disable default Apache virtual host
  apache2_module:
    name: default
    state: disabled

- name: Enable configured Apache virtual host
  apache2_module:
    name: your_domain
    state: enabled

- name: Copy web page content
  copy:
    content: |
      <html>
      <head>
      <title>Hello World</title>
      </head>
      <body>
      <h1>Hello World!</h1>
      </body>
      </html>
    dest: /var/www/html/index.html

handlers:
  - name: restart Apache
    service:
      name: apache2
      state: restarted
```
</details>

## 4. Create the Apache virtual host template:
- Inside the **`roles/webserver`** directory, create a file called templates/vhost.conf.j2.
- Add the following content to **`vhost.conf.j2`**:
<details>
  <summary>apache</summary>

    <VirtualHost *:80>
    ServerName your_domain
    Redirect permanent / https://your_domain/
    </VirtualHost>

        <IfModule mod_ssl.c>
    <VirtualHost *:443>
    ServerName your_domain
    SSLEngine on
    SSLCertificateFile /etc/ssl/certs/selfsigned.crt
    SSLCertificateKeyFile /etc/ssl/private/selfsigned.key
    ProxyPreserveHost On
    ProxyRequests Off
    ProxyPass /
    </VirtualHost>
    </IfModule>

</details>

These steps will help you set up the Ansible project, create the playbook, define the role, and create the necessary configuration files for configuring and deploying a web server.