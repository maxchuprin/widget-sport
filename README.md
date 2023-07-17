# Widget-Sports
> Этот сервис отвечает за получение анонсов ближайших событий, а так же счёт матча в режиме реального времени

#### Environment
- **{rapidApiKey}** - ключ Rapid API
- **{footballHost}** - HOST Rapid API футбола
- **{hockeyHost}** - HOST Rapid API хокея
- **{basketballHost}** - HOST Rapid API баскетбола
- **{tennisHost}** - HOST API тенниса
- **{mmaHost}** - HOST API MMA
- **{timeZone}** - таймзона получаемых данных(детали по остальным часовым поясам уточнять)
- **{plusDay}** - дата получения анонсов матчей. Необходимо указать цифру, через сколько дней. (по умолчанию цифра 1, соотвественно анонс на следующий день)
- **{cacheMinutes}** - время кэширования запросов
- **{footballLeague}** - Список футбольных лиг для фильтра
- **{footballTeam}** - Список футбольных команд для фильтра
- **{hockeyLeague}** - Список хоккейных лиг для фильтра
- **{basketballLeague}** - Список баскетбольных лиг для фильтра
- **{tennisLeague}** - Список теннисных лиг для фильтра


#### Тип возвращаемых данных
- **JSON**

#### endPoint
- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/football/fixtures-in-progress}** счёт футбола в реальном времени
- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/football/fixtures-by-date}** анонсы футбола
  https://rapidapi.com/api-sports/api/api-football/

- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/hockey/fixtures-in-progress}** счёт хоккея в реальном времени
- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/hockey/fixtures-by-date}** анонсы хоккея
  https://rapidapi.com/api-sports/api/api-hockey/

- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/basketball/fixtures-in-progress}** счёт баскетобола в реальном времени
- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/basketball/fixtures-by-date}** анонсы баскетбола
  https://rapidapi.com/api-sports/api/api-basketball/

- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/tennis/fixtures-in-progress}** счёт тениса в реальном времени
- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/tennis/fixtures-by-date}** анонсы тениса
  https://rapidapi.com/tipsters/api/sportscore1/

- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/mma/results}** результаты боев
- {SERVER_ADDRESS}:{SERVER_PORT}/**widget-sport**/mma/fixtures-by-date}** анонсы боев
  https://rapidapi.com/tipsters/api/flashlive-sports

#### Методы
- **GET**