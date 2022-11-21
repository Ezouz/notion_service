function send(from, to, mailjet, object, text) {
  return new Promise((done, err) => {
    mailjet.post("send", { version: "v3.1" }).request({
      Messages: [
        {
          From: {
            Email: from,
            Name: "Admin",
          },
          To: [
            {
              Email: to,
              Name: "Admin",
            },
          ],
          Subject: object,
          TextPart: "What is text part",
          HTMLPart: text,
          CustomID: "coucou",
        },
      ],
    });
  });
}

module.exports = { send };
