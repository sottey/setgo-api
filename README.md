
![Logo](https://seanottey.elzs.net/setgo-api.png)

An unofficial, unauthorized, unknown, unapproved caching API for [SetApp](https://setapp.com) data

## Features

- Programmatic access to SetApp data
- Complete app data including icons, links, release date and more
- Caching to minimize impact on SetApp servers
- Full text search of app names and descriptions
- Category listing
- Subcategory id and name listing
- Apps by category
## API REFERENCE

#### SHOW HELP

```http
  GET /
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `none` | `none` | `none` |

*Returns API information*


#### GET ALL APPS

```http
  GET /apps
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| none     | none | none |

*Returns all apps*


#### GET APP INFORMATION BY ID
```http
  GET /apps/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`     | `string` | `ID of app in collection` |

*Returns app information for app with specified id*

#### GET ALL CATEGORY NAMES
```http
  GET /cats
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `none`     | `none` | `none` |

*Returns all category names*

#### GET ALL APPS IN CATEGORY
```http
  GET /cats/{category}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category`     | `string` | `Name of category` |

*Returns all apps in specified category*

#### GET SUBCATEGORY IDS AND NAMES
```http
  GET /subcats
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `none`     | `none` | `none` |

*Returns all subcategories*

#### SEARCH APPS
```http
  GET /search/{query}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `query`     | `string` | `Term to search for` |

*Return all items which have the query provided in their name or description*



## TECH STACK

- [Go](https://go.dev)
- [SetApp data](https://setapp.com)


## DEPLOYMENT

To deploy this project get the release executable for your OS or clone the source and build 

```bash
  git clone https://www.github.com/sottey/setgo-api
  sudo go build setgo-api.go
```

## .ENV FILE
To run, edit the **./.env** file and set the following values:
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| SETAPP_URL| string | SetApp apps page (probably https://setapp.com/apps) |
| CACHE_FILE | string | File location of cached json data |
| SUBCAT_FILE | string | Id and name mapping json |
| HELP_TEMPLATE | string | File location of HTML template for help (/)|
| SERVER_URL | string | URL and port to server the data |
| FAVICON | string | File location of favicon | 

## RUNNING THE SERVER
```bash
  ./setgo-api
```

Once running, you will be able to access the help page at the url specified in the .env file


## LICENSE

[MIT](https://choosealicense.com/licenses/mit/)


## AUTHORS

- [@sottey](https://www.github.com/sottey)

