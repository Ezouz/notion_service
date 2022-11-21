const { Client, LogLevel } = require("@notionhq/client");
const dotenv = require("dotenv");
const Mailjet = require("node-mailjet");
const page = require("./src/page");
const checks = require("./src/checks");
const email = require("./src/emails");

dotenv.config({ path: "../.env" });

const mailjet = Mailjet.apiConnect(
  process.env.MAILJET_API_KEY,
  process.env.MAILJET_API_SECRET
);

// Initializing a client
const notion = new Client({
  auth: process.env.NOTION_TOKEN,
  // logLevel: LogLevel.DEBUG,
});

(async () => {
  main();
})();

let tasksInDbz = [];

async function main() {
  const dbList = [
    {
      db: process.env.DB_ORGABOCAL,
      name: "Event / Organization",
      check: ["status"],
      notify: process.env.ZOUZ_MAIL,
    },
    {
      db: process.env.DB_DOC,
      name: "Documentation",
      check: ["status"],
      notify: process.env.NOTION_MAIL,
    },
    {
      db: process.env.DB_SERVICE,
      name: "Services",
      check: ["status"],
      notify: process.env.NOTION_MAIL,
    },
    {
      db: process.env.DB_TEST,
      name: "TEST",
      check: ["status"],
      notify: process.env.ZOUZ_MAIL,
    },
  ];
  let dbs = [];

  // get db
  for (const [key, value] of Object.entries(dbList)) {
    const content = await page.getContentFromDatabase(notion, value.db);
    dbs.push({
      dabatase_id: value.db,
      dabatase_name: value.name,
      check: value.check,
      notify: value.notify,
      content,
    });
  }

  // check content
  let checked = [];
  for (const [id, value] of Object.entries(dbs)) {
    const list = await checks.mapChanges(
      value,
      tasksInDbz.filter((e) => e.dabatase_id === value.dabatase_id)[0]
    );
    checked.push({
      dabatase_id: value.dabatase_id,
      name: value.dabatase_name,
      notify: value.notify,
      list,
    });
  }
  // post email or discord message with change
  for (const [index, database] of Object.entries(checked)) {
    if (database.list.length > 0) {
      for (const [idx, elem] of Object.entries(database.list)) {
        for (const [idx, change] of Object.entries(elem)) {
          // console.log(
          //   `to : ${database.notify}`,
          //   `${change.action} in database ${database.name} row : ${change.row}`,
          //   `line : ${change.row}
          //   property : ${change.props}
          //   action: ${change.action}
          //   from ${change.old} to ${change.new}`
          // );
          email.send(
            process.env.NOTION_MAIL,
            database.notify,
            mailjet,
            `${change.action} in database ${change.dabatase_name} row : ${change.row}`,
            `line : ${change.row}
              property : ${change.props}
              action: ${change.action}
              from ${change.old} to ${change.new}
              page url : ${change.url}
              `
          );
        }
      }
    }
  }

  // keep content
  tasksInDbz = dbs;

  setTimeout(await main(), 5000);
}
