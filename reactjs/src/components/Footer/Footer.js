import { Link } from "react-router-dom";
import styled from "styled-components";

const Parent = styled.div`
  padding: 30px 20px;
`;

const Footer = () => {
  return (
    <Parent>
      <Link to="/privacy">Privacy policy</Link>
    </Parent>
  );
};

export default Footer;
