import styled from "styled-components";
import { Link } from "react-router-dom";

const Parent = styled.div`
  padding: 20px;
`;

const PrivacyPolicy = () => {
  return (
    <Parent>
      <Link to="/">Back</Link>

      <h2>Privacy Policy</h2>

      <p>At our service, we take your privacy seriously. Below are some key points regarding how we handle your information:</p>

    <ul>
        <li>Any information you input remains within your browser, except for the domains you provide in the input text. These domains are necessary for processing your request and are securely stored in our database.</li>
        <li>The stored domains are used solely for the tasks you requested and are never shared, sold, or utilized for any undisclosed purposes.</li>
        <li>We are committed to maintaining the confidentiality and security of your data. Your trust is important to us, and we will never misuse your information.</li>
        <li>If you wish to delete the stored data from our database, please contact us at <a href="mailto:admin@ginam.us">admin@ginam.us</a>.</li>
    </ul>

    <p>If you have any concerns or questions regarding our privacy practices, please don't hesitate to contact us.</p>
    </Parent>
  );
};

export default PrivacyPolicy;
