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
  const [isLoaded, setIsLoaded] = useState(false);
  const [progress, setProgress] = useState(0);
  const [mermaidChart, setMermaidChart] = useState("");
  const [domains, setDomains] = useState("");
  const [jobId, setJobId] = useState("");

  useEffect(() => {
    let timerId = setInterval(() => {
      if (jobId) {
        (async () => {
          const r = await api.get(`/${jobId}`);
          setProgress(r.progress);
          if (r.completed) {
            setJobId(""); // stop checker
            setIsLoading(false);
            setIsLoaded(true);
          }
          const m = jobUtil.dataToMermaidGraphDefinition(r);
          console.log(m);
          setMermaidChart(m);
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
    setIsLoaded(false);
    setProgress(0);

    // parse domains
    const parts = domains.trim()
      .split("\n")
      .filter(v => v.trim() !== "");

    const {jobId: rJobId} = await api.post("/", {
      domains: parts,
    });
    console.log({rJobId});
    setJobId(rJobId);
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
      {isLoading && (
        <Progress progress
                  success={isLoaded}
                  size='small'
                  percent={progress}>Building diagram 🔨 ...</Progress>
      )}

      <Mermaid chart={mermaidChart} />

    </Parent>
  );
};

export default Facade;
