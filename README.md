# osu-dashboard

### Requirements
- docker
- go 1.23+
- nodejs 20+
- npm

### Run locally

- Create `.env` file inside `backend` dir and fill with your info

```shell
OSU_API_CLIENT_ID=
OSU_API_CLIENT_SECRET=
POSTGRES_PASSWORD=
```
- Start backend and frontend

```shell
# backend
cd backend
make up
# frontend
cd frontend
npm run dev
```