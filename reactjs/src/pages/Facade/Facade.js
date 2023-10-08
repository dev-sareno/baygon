import styled from "styled-components";
import {Button, Form, Progress, TextArea} from 'semantic-ui-react'
import React, {useEffect, useState} from "react";
import mermaid from "mermaid";
import Mermaid from "../../components/Mermaid/Mermaid";

mermaid.initialize({startOnLoad: true});

const Parent = styled.div`
  padding: 20px;
`;

const Facade = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [progress, setProgress] = useState(44);
  const [isLoaded, setIsLoaded] = useState(false);
  const [mermaidChart, setMermaidChart] = useState("");

  const lookupClickHandler = async () => {
    setIsLoading(true);

    const graphDefinition = `
    graph LR
      A --- B
      B-->C[fa:fa-ban forbidden]
      B-->D(fa:fa-spinner);
    `;
    setMermaidChart(graphDefinition);
  };

  return (
    <Parent>
      <Form>
        <div>Target DNS:</div>
        <TextArea placeholder={"www.google.com\nwww.youtube.com"} style={{minHeight: 100}}/>
        <Button primary
                loading={isLoading}
                style={{marginTop: 10}}
                onClick={lookupClickHandler}>Lookup</Button>
      </Form>

      <br/>
      <br/>
      <Progress progress
                size='small'
                percent={progress}>Building diagram 🔨 ...</Progress>

      <Mermaid chart={mermaidChart} />

    </Parent>
  );
};

export default Facade;
