import sha1 from "js-sha1";

const dataToMermaidGraphDefinition = (data) => {
  const lines = [];
  for (let i=0; i < data.data.length; i++) {
    const v = data.data[i];

    const domain = v.domain;
    const cname = v.cname;
    const a = v.a;

    // remove empty values
    const clean = [domain, cname, a].filter(j => j.trim() !== "");

    // replace new line with <br>
    const formatted = clean.map(j => j.replaceAll("\n", "<br>"));

    // assign keys (base64)
    const withKeys = formatted.map(j => {
      const key = sha1.hex(j).substring(0, 6);
      return `${key}[${j}]`;
    });

    // join
    const mermaid = withKeys.join(" --> ");
    lines.push(mermaid);
  }

  return [
    "flowchart LR",
    lines.map(j => "    " + j).join("\n"),
  ].join("\n");
};

export default {
  dataToMermaidGraphDefinition,
}
