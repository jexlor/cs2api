## Cs2Api   ![Go Version](https://img.shields.io/badge/Go-1.23.3-blue) [![PkgGoDev](https://pkg.go.dev/badge/github.com/jexlor/cs2api)](https://pkg.go.dev/github.com/jexlor/cs2api)

# <img alt="cs2api" src="https://github.com/user-attachments/assets/a1dce9fe-507c-410d-9e4d-142d8b4cef13" width="220" />

What's Cs2Api?
---------------------------
<strong>cs2api<strong> repo contains database filled with all the information about skins from cs2 (game). all that is accessible with 
API which is written with <strong>Go + Gin</strong> framework. for DB I'm using <strong>PostgreSql<strong>. Api can work for any type of project which needs db+api to serve info about thousands of skins.
(for now, api is done but database is empty. it has only some rows for testing. there is still work to do.. i'm working on web scraper which should take all info from various 
sites, convert it in neccessary form and add to database. since i want all information to be correct and updated weekly, it takes time, i'm also working to make copy of database with docker volumes so you can iterate over db.
if you want to test api now, consider adding some rows to your local postgre database with credentials in <strong>.example.env</strong> file.)

Getting started with Cs2api.
----------------------------
```bash
git clone https://github.com/jexlor/cs2api.git
```

First step (copy .env file in repo root)
----------------------------
```bash
cp .example.env .env
```
Run Docker:
----------------------------
```bash
docker compose up --build
```
Access Api welcome page (localhost:8080/cs2api by default):
---------------------------
![croped](https://github.com/user-attachments/assets/2bbce750-0db8-4d1c-a2bd-76cfae1d9efb)

