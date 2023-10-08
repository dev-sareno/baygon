import styled from "styled-components";
import {Form, TextArea} from 'semantic-ui-react'

const Parent = styled.div`
  padding: 20px;
`;

const Facade = () => {
  return (
    <Parent>
      <Form>
        <TextArea placeholder='Tell us more' style={{minHeight: 100}}/>
      </Form>
    </Parent>
  );
};

export default Facade;