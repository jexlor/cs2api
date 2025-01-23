## Cs2Api  [![PkgGoDev](https://pkg.go.dev/badge/github.com/jexlor/cs2api)](https://pkg.go.dev/github.com/jexlor/cs2api)

# <img alt="cs2api" src="https://github.com/user-attachments/assets/a1dce9fe-507c-410d-9e4d-142d8b4cef13" width="220" />

What's Cs2Api?
---------------------------
<strong>cs2api</strong> repo contains database filled with all the information about skins from cs2 (game). all that is accessible with 
Api which is written with <strong>Go + Gin</strong> framework. for DB Api has <strong>PostgreSql</strong>. Api can work for any type of project which needs db + api to serve info about thousands of skins.

Where Cs2Api gets information about skins?
-----------------------------------------------
<strong>cs2api</strong> database is filled with real data from real sites. api uses scraping tool which will be released on that account once it's 100% ready and useful, basically data on that api is the same you see on major sites. scraper automatically updates prices of that data every week (still in development)

Total skins: <strong>603</strong> <br>
Last updated prices & collections: <strong>23.1.25</strong>

Additional features which you can enable (disabled by default): <strong>PgAdmin</strong>, <strong>DB Migrations</strong>, <strong>JWT</strong>, <strong>Rate limiting</strong>

Structure of skins
----------------------------
```json
{
    "id": 1,
    "name": "test_name",
    "weapon": "AK-47"
    "rarity": "covert",
    "collection": "bravo",
    "price": "$100-$150",
    "stattrack_price": "$500-$750",
    "url": "example.com/skin_image_url"
}
```

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
Run Docker container:
----------------------------
```bash
docker compose up --build
```
Access Api welcome page (on default cs2api endpoint):
---------------------------
![cs2api](https://github.com/user-attachments/assets/054f00f3-aa2f-4b69-a2fe-dc15fdcc0c68)


