async function getContentFromDatabase(notion, database_id, cursor) {
  let content = [];
  const res = await getContentFromDatabaseAtPage(notion, database_id, cursor);
  if (res.next_cursor) {
    content = [...res.content];
    await getContentFromDatabase(notion, database_id, cursor);
  } else {
    content = [...res.content];
  }

  return content;
}

async function getContentFromDatabaseAtPage(notion, database_id, cursor) {
  let request_payload = "";
  //Create the request payload based on the presense of a start_cursor
  if (cursor == undefined) {
    request_payload = {
      path: "databases/" + database_id + "/query",
      method: "POST",
    };
  } else {
    request_payload = {
      path: "databases/" + database_id + "/query",
      method: "POST",
      body: {
        start_cursor: cursor,
      },
    };
  }

  //While there are more pages left in the query, get pages from the database.
  let current_pages;
  try {
    current_pages = await notion.request(request_payload);
  } catch (error) {
    console.error(error);
  }

  return {
    content: current_pages.results,
    next_cursor: current_pages.has_more ? current_pages.next_cursor : null,
  };
}

module.exports = { getContentFromDatabase, getContentFromDatabaseAtPage };
