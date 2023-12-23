[![Main-Docker](https://github.com/aceberg/exercisediary/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/exercisediary/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/exercisediary)](https://goreportcard.com/report/github.com/aceberg/exercisediary)
[![Maintainability](https://api.codeclimate.com/v1/badges/e8f67994120fc7936aeb/maintainability)](https://codeclimate.com/github/aceberg/exercisediary/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/exercisediary)

<h1><a href="https://github.com/aceberg/exercisediary">
    <img src="https://raw.githubusercontent.com/aceberg/exercisediary/main/assets/logo.png" width="35" />
</a>Exercise Diary</h1>

Light and easy workout diary

- [Quick start](https://github.com/aceberg/exercisediary#quick-start)
- [Config](https://github.com/aceberg/exercisediary#config)
- [Options](https://github.com/aceberg/exercisediary#options)
- [Thanks](https://github.com/aceberg/exercisediary#thanks)


![Screenshot](https://raw.githubusercontent.com/aceberg/exercisediary/main/assets/Screenshot%202023-04-02%20at%2022-27-40%20Resource%20Diary.png)

## Quick start

```sh
docker run --name exercisediary \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/ExerciseDiary:/data/ExerciseDiary \
-p 8851:8851 \
aceberg/exercisediary
```
<!-- Or use [docker-compose.yml](docker-compose.yml) -->


## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8851 |
| THEME | Any theme name from https://bootswatch.com in lowcase | grass |
| COLOR | Background color: light or dark | light |
| HEATCOLOR | HeatMap color | #03a70c |
| TZ | Set your timezone for correct time | "" |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -c | Path to config dir | /data/ExerciseDiary |  

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/exercisediary/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Flaticon](https://www.flaticon.com/icons/)