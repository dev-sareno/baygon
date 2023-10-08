import styled from "styled-components";

const Parent = styled.div`
  border-bottom: 1px solid #EAECEE;
  padding: 16px 8px;
  font-size: 24px;
`;

const Title = styled.div`
  color: #2471A3;
`;

const Header = () => {
  return (
    <Parent>
      <Title>Ginamus - DNS Lookup Tool</Title>
    </Parent>
  );
};

export default Header;
