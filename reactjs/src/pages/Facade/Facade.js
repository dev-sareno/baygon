import styled from "styled-components";
import {Button, Form, Progress, TextArea} from 'semantic-ui-react'
import React, {useEffect, useState} from "react";
import mermaid from "mermaid";
import Mermaid from "../../components/Mermaid/Mermaid";
import api from "../../api/api";
import jobUtil from "../../util/job-util";

mermaid.initialize({startOnLoad: true});

const Parent = styled.div`
  padding: 20px;
`;

const Facade = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [progress, setProgress] = useState(44);
  const [isLoaded, setIsLoaded] = useState(false);
  const [mermaidChart, setMermaidChart] = useState("");
  const [domains, setDomains] = useState("");
  const [jobId, setJobId] = useState("");

  useEffect(() => {
    let timerId = setInterval(() => {
      if (jobId) {
        (async () => {
          const r = await api.get(`/${jobId}`);
          if (r.completed) {
            setJobId(""); // stop checker
          }
          const m = jobUtil.dataToMermaidGraphDefinition(r);
          console.log(m);
        })();
      }
    }, 5000);
    return () => {
      clearInterval(timerId);
    }
  }, [jobId]);

  const lookupClickHandler = async () => {
    if (isLoading) {
      return;
    }
    setIsLoading(true);

    // parse domains
    const parts = domains.trim()
      .split("\n")
      .filter(v => v.trim() !== "");
    console.log({parts})

    const {jobId: rJobId} = await api.post("/", {
      domains: parts,
    });
    console.log({rJobId});
    setJobId(rJobId);

    //
    // const graphDefinition = `
    // graph LR
    //   A --- B
    //   B-->C[fa:fa-ban forbidden]
    //   B-->D(fa:fa-spinner);
    // `;
    // setMermaidChart(graphDefinition);
  };

  const domainsChangeHandler = (e) => {
    setDomains(e.target.value);
  };

  return (
    <Parent>
      <Form>
        <div>Target DNS:</div>
        <TextArea style={{minHeight: 100}}
                  onInput={domainsChangeHandler}
                  placeholder={"www.google.com\nwww.youtube.com"}/>
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
