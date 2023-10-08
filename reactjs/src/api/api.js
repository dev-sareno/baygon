const BASE_URL = process.env.REACT_APP_BACKEND_URL;

const get = async (path) => {
  const r = await fetch(BASE_URL + path, {
    method: "GET",
    headers: {
      // "Accept": "application/json",
    }
  });
  const getResponse = await r.json();
  console.log({getResponse});
  return getResponse;
}

const post = async (path, body) => {
  const r = await fetch(BASE_URL + path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      // "Accept": "application/json",
    },
    body: JSON.stringify(body),
  });
  const postResponse = await r.json();
  console.log({postResponse});
  return postResponse;
};

export default {
  get,
  post
}
