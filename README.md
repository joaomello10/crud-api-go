


# crud api go

![GitHub repo size](https://img.shields.io/github/repo-size/joaomello10/crud-api-go?style=for-the-badge)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://camo.githubusercontent.com/544022edf8369d944e68802fc043b0268484709e334d23db2882590aeae296cb/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f706f7374677265732d2532333331363139322e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d706f737467726573716c266c6f676f436f6c6f723d7768697465)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)


## Simple Crud to manage notes made with go


<img src="image.png" alt="Example image">

## 🚀 Installing

### To install, follow these steps:

Linux and macOS:

```
git clone https://github.com/joaomello10/crud-api-go.git
```

## ☕ Using API

### To Start API, follow these steps:

```
docker-compose up -d
```

### then postgres start, run:

```
go run cmd/api/main.go
```

## The api will run default in `http://localhost:8080`

## Routes

### Get all notes

**GET** `/notes`

#### Response
```json
[
  {
    "ID": 3,
    "Title": "Foo",
    "Body": "Bar",
    "created_at": "2025-03-27T18:48:49.836731-03:00",
    "updated_at": "2025-03-27T18:48:55.408299-03:00"
  },
  {
    "ID": 4,
    "Title": "Bar",
    "Body": "Qx",
    "created_at": "2025-03-27T18:48:49.836731-03:00",
    "updated_at": "2025-03-27T18:48:55.408299-03:00"
  }
]
```
### Get especified notes

**GET** `/note/:id`

#### Response `/note/3`
```json
[
  {
    "ID": 3,
    "Title": "Foo",
    "Body": "Bar",
    "created_at": "2025-03-27T18:48:49.836731-03:00",
    "updated_at": "2025-03-27T18:48:55.408299-03:00"
  },
]
```
### Create Note

**POST** `/new`

#### Body 
```json
[
  {
    "Title": "Foo",
    "Body": "Bar",
  },
]
```
### Update note

**PUT** `/note/:id`

#### Body
```json
[
 {
    "Title": "Foo 2",
    "Body": "Bar 2",
  },
]
```

### Delete note

**DELETE** `/note/:id`

#### Response `/note/3`
```json
[
 {"message": "Post deleted successfully"}
]
```

## 📝 License

This project is under license. See the [LICENSE](LICENSE.md) file for more details.
