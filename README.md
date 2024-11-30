## Cs2Api  [![PkgGoDev](https://pkg.go.dev/badge/github.com/jexlor/cs2api)](https://pkg.go.dev/github.com/jexlor/cs2api)

# <img alt="cs2api" src="https://github.com/user-attachments/assets/a1dce9fe-507c-410d-9e4d-142d8b4cef13" width="220" />

What's Cs2Api?
---------------------------
<strong>cs2api</strong> repo contains database filled with all the information about skins from cs2 (game). all that is accessible with 
API which is written with <strong>Go + Gin</strong> framework. for DB I'm using <strong>PostgreSql</strong>. Api can work for any type of project which needs db+api to serve info about thousands of skins.
(Volume of database will be added to repo as soon as it's ready)

Cases and collections supported:

<strong>Gallery case, </strong>
<strong>Kilowatt Case, </strong>
<strong>Revolution Case, </strong> 
<strong>Recoil Case, </strong>
<strong>Dreams and Nightmares Case, </strong>

Total skins: <strong>66</strong> <br>
Last updated prices & collections: <strong>30.11.24</strong>

Structure of skins
----------------------------
```json
{
    "id": 1
    "name": "test_name",
    "rarity": "covert",
    "collection": "bravo",
    "quality": "Factory New",
    "price": "$ 100",
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
Run Docker:
----------------------------
```bash
docker compose up --build
```
Access Api welcome page (localhost:8080/cs2api by default):
---------------------------
![cs2api](https://github.com/user-attachments/assets/054f00f3-aa2f-4b69-a2fe-dc15fdcc0c68)


