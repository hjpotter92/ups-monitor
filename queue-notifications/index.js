const AWS = require("aws-sdk");

const snsTopicArn = process.env.SNS_TOPIC_ARN;

exports.handler = async function (event) {
  const body = JSON.parse(event.body);
  const message = {
    Message: JSON.stringify(body),
    TopicArn: snsTopicArn,
  };
  return new AWS.SNS({
    apiVersion: "2010-03-31",
  })
    .publish(message)
    .promise();
};
