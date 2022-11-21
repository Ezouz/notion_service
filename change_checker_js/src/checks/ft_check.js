// checks if status are newly set as complete or equivalents listed in statusAlert
// not working on reopening
function status(newList, oldList) {
  const statusAlert = ["down", "deprecated", "maintenance", "done"];
  const pages = [];

  // check page
  for (const [page_idx, page] of Object.entries(newList)) {
    const changes = [];
    for (const [key, props] of Object.entries(page.properties)) {
      // check if it as one or more status properties
      // check if properties values are in statusAlert
      if (
        props.type === "status" &&
        statusAlert.filter(
          (stat) => stat === props.status.name.toLowerCase()
        )[0]
      ) {
        // if yes check it's equivalent in oldList
        if (oldList) {
          oldPage = oldList.filter((oldP) => oldP.id === page.id)[0];
          for (const [key, oldProps] of Object.entries(oldPage.properties)) {
            if (
              oldProps.type === "status" &&
              props?.id === oldProps?.id &&
              props.status?.name?.toLowerCase() !==
                oldProps.status?.name?.toLowerCase()
            ) {
              changes.push({
                props: key,
                action: "change",
                row: "",
                new: props.status?.name,
                old: oldProps.status?.name,
                url: page.url,
              });
            }
          }
        }
      }
      if (props.type === "title" && changes.length > 0) {
        changes.map((e) => (e.row = props?.title[0]?.text?.content));
      }
    }
    if (changes.length > 0) {
      pages.push(changes);
    }
  }
  return pages;
}

module.exports = { status };
