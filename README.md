## URL Shortener Demo

### 使用方式
    # Upload URL API
    curl -X POST -H "Content-Type:application/json" http://localhost:8080/api/v1/urls -d '{
        "url": "<original_url>",
        "expireAt": "2021-02-08T09:20:41Z"
    }'
    # Response
    {
        "id": "<url_id>",
        "shortUrl": "http://localhost/<url_id>"
    }
    # ------------------
    
    # Redirect URL API
    curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL
    
#### 安裝
1. 安裝 [Docker](https://docs.docker.com/engine/install/)
2. 安裝 [Docker Compose](https://docs.docker.com/compose/install/)
3. 執行 `docker-compose build --no-cache`

### 啟動
執行 `docker-compose up`

### database
由於想要使用auto increment來自動產生url_id，故選擇使用關連式資料庫  
而決定使用mysql的原因則是因為對於mysql較為熟悉

### 套件說明
##### gin
    github上star跟contributors最多的golang web framework, 文件資料跟各種支援也最豐富
##### logrus
    功能強大的logger, 還有很多plugin可以串接elk、slack等服務，也可以根據界面開發自己的plugin
##### gorm and gorm/mysql
    golang orm, 因為題目沒有複雜的資料庫查詢，故使用orm縮短開發時間
##### go-cache
    由於資料內容不會變動，故使用本地端cache儲存db查詢結果增加api效能
##### go-sqlmock
    在測試時對db進行mock的工具
##### golang/mock
    在測試時對interface進行mock的工具
##### goconvey
    golang的測試框架
##### go-hashids
    用來替url的id進行hash產生url_id
    
### script 說明
##### Dockerfile
    用來製作URLShortenerDemo的docker image
##### docker-compose.yml
    用來啟動mysql跟URLShortenerDemo的docker container的script
##### wait-for-it.sh
    在docker-compose.yml中用來讓URLShortenerDemo等待mysql啟動完再執行指令的script
##### start.sh
    在wait-for-it.sh結束後執行的指令，會先進行db的migrate後再啟動URLShortenerDemo
##### coverage.sh
    計算coverage的script，結果儲存在coverage內