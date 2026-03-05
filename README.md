# Email-Sender-Cli

## Requirements

- [x] Uses .env file for configurations
- [x] Read messages from table `outbox`
- [x] Connect to email service (provided in the .env)
- [x] Send messages with attribute `0` ( column `sent`)
- [x] Modifiy sent messages state in the database
- [x] Send messages concurrently
- [ ] Write test for race conditions

## Implementation

- the `outbox` table

```SQL
CREATE TABLE outbox (
    id INT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    subject VARCHAR(255) ,
    body TEXT ,
    sent TINYINT DEFAULT 0
);
insert into outbox (address, subject, body, sent) values ('hbims0@dailymotion.com', 'bibendum imperdiet', 'Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo.', 1);
```

- Least SMTP commands (tell a client or server what to do and how to handle any accompanying data.)
  - `FROM`
  - `TO`
  - `Subject`

## Modules Needed

- Mariadb driver "github.com/go-sql-driver/mysql"
- .env parser "github.com/joho/godotenv"
