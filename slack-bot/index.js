const webApi = require('@slack/web-api');
const token = process.env.SLACK_TOKEN;
const channel = process.env.SLACK_CHANNEL;

const client = new webApi.WebClient(token);

exports.handler = async() => {
    const res = await client.chat.postMessage({
        channel,
        text: "Powercut at my place.",
        as_user: true
    });
    console.log(res);
};
