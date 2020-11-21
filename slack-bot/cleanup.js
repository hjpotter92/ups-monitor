const webApi = require('@slack/web-api');
const token = process.env.SLACK_TOKEN;
const channel = process.env.SLACK_CHANNEL;

const client = new webApi.WebClient(token);

(async () => {
    const res = await client.chat.delete({
        channel,
        ts: '1605854808.001400',
        as_user: true
    });
    console.log(res);
})();
