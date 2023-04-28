# LIMe

LIMe is a simple LINE Messaging API Webhook receiver.

## Usage

### 1. Create a new LINE Bot

Create a new LINE Bot and get the Channel Secret and Channel Access Token.

### 2. Implement EventHandler

Implement a handler for each [event type](https://github.com/line/line-bot-sdk-go/blob/master/linebot/event.go) you wish to hook.

A simple sample is presented below.  
- [Echo message handler](example/echo.go)

### 3. Deploy LIMe with your EventHandler

Deploy LIMe to your server.

#### Environment Variables

Set the following environment variables.

| Name                          | Description                                                                                                    | Required | Default                    |
|-------------------------------|----------------------------------------------------------------------------------------------------------------|----------|----------------------------|
| `PORT`                        | Port number to listen                                                                                          | No       | `8080`                     |
| `HandlePath`                  | Path to handle webhook requests                                                                                | No       | `/callback`                |
| `CHANNEL_SECRET`              | Channel Secret of your LINE Bot                                                                                | Yes      |                            |
| `CHANNEL_TOKEN`               | Channel Access Token of your LINE Bot                                                                          | Yes      |                            |
| `READ_TIMEOUT_SEC`            | Timeout to read request                                                                                        | No       | `5`                        |
| `WRITE_TIMEOUT_SEC`           | Timeout to write response                                                                                      | No       | `10`                       |
| `IDLE_TIMEOUT_SEC`            | Timeout to keep connection                                                                                     | No       | `120`                      |
| `EVENT_TIMEOUT_SEC`           | Set context timeout time per event, noting that multiple events may be requested together in a single webhook. | No       | `10`                       |
| `LINE_API_ENDPOINT_BASE`      | LINE Messaging API endpoint base URL                                                                           | No       | `https://api.line.me`      |
| `LINE_API_ENDPOINT_BASE_DATA` | LINE Messaging API endpoint base URL for data                                                                  | No       | `https://api-data.line.me` |
| `ENABLE_RETURN_ERROR_CODE`    | Enable to return error code                                                                                    | No       | `false`                    |

##### Note

`ENABLE_RETURN_ERROR_CODE`

If True, returns a code other than 200 when an error occurs in EventHandler.
If Webhook resend is enabled in LINE Messaging API, the same event will be resent.  
Please implement the following in event handlers.

- Do not re-execute events with WebhookEventIDs that have already been processed.
- Or, the processing in the handler should idempotent.

## 4. Operation check

Check the operation by sending a message to the LINE Bot.
