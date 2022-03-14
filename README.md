
# <p  align="center" style="font-family:courier;font-size:130%" size=212px> LEMONLOG - easy logging solution </p> 


[![Generic badge](https://img.shields.io/badge/LICENSE-MIT-orange.svg)](LICENSE)

<p align="center">
  <img src="https://cdn-icons.flaticon.com/png/512/2732/premium/2732012.png?token=exp=1647252081~hmac=c496155513b143679c9a5a68c6a99b66" alt="logo"/>
</p>

Lemonlog - is a loghub, that can be easyly deployed in a docker container and connected to a docker microservices infrastructure. 

It is based on mongo image, and supports gelf udp input.

It is lightweight, easy to learn and working with lots of different docker images straight out of the box.

LemonLog can be configured throw env-variables in docker-compose file, so that there is no need in manual tuning.

LemonLog supports logs from many different technologies and automatically prepares dashboards and basic analytics.

Also it is easy to set email notifications using lemon-log, additional http requests by triggers can be added aswell.

# Search

You can store input data in json format, which can then bequeried in underlying mongo database using simlpe for human understanding queries. Data, that can be queried is automatically filtered by tags and levels and can contain additional query field to search by other parameters.

# Visualize

You can set custom made visualizations in lemonlog, that are powered by [fl_chart](https://pub.dev/packages/fl_chart) flutter package. Also a lot of technologies are automatically recognized by logger, and you don't have to manually setup the chart for them - they are working out of the box.

# Alert

Lemonlog supports executing http requests on alters, which allows to connect different forms of alertions:
- email
- sms-services
- messangers
- api actions

# Deploy

All settings can be configured as environment variables on app launch, so that there is no need for manual clicks inside the logger itself.

