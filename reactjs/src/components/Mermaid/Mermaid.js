import React, {useEffect, useState} from "react";
import mermaid from "mermaid";

mermaid.initialize({startOnLoad: false});

const emptyGraph = "graph LR";

const Mermaid = ({chart}) => {
  const [graphDefinition, setGraphDefinition] = useState(emptyGraph);

  useEffect(() => {
    if (chart) {
      setGraphDefinition(chart);
    }
  }, [chart]);

  useEffect(() => {
    (async () => {
      const { svg } = await mermaid.render('outputGraph', graphDefinition);
      const outputElement = document.querySelector("#mermaid");
      console.log({outputElement});
      if (outputElement) {
        outputElement.innerHTML = svg;
      }
    })();
  }, [graphDefinition]);

  return (
    <div id="mermaid"></div>
  );
};

export default Mermaid;
