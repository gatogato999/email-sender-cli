# Email-Sender-Cli

## Requirements

- [x] Uses .env file for configurations
- [x] read messages from table `outbox`
- [ ] Connect to email service (provided in the .env)
- [ ] send messages with attribute `false` ( column `sent`)

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

## Modules Needed

- Mariadb driver "github.com/go-sql-driver/mysql"
- .env parser "github.com/joho/godotenv"
