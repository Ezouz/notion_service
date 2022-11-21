const ft_checks = require("./ft_check");

function mapChanges(newContent, oldContent) {
  let checks = [];
  for (const check of newContent.check) {
    switch (check) {
      case "status":
        checks = [...ft_checks.status(newContent.content, oldContent?.content)];
        break;
      default:
        break;
    }
  }
  return checks;
}

module.exports = { mapChanges };
