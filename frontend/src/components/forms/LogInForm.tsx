import React, { useState } from "react";
import { Redirect } from "react-router-dom";
import { Pane, Button, Heading, TextInputField } from "evergreen-ui";
import API from "./../../Api";
import * as types from "./../../types";

type Props = { getSessionData: () => void };
const LogInForm = (props: Props) => {
  const emptyMessages: types.ValidationMessages = {
    email: "",
    password: "",
  };
  const [validationMessages, setValidationMessages] = useState(emptyMessages);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [shouldRedirect, setShouldRedirect] = useState(false);

  const submitForm = () => {
    API.post("/users/login", {email: email, password: password})
      .then( response => {
        if (response.data && response.data.status === "ok") {
          setValidationMessages(emptyMessages);
          props.getSessionData();
          setShouldRedirect(true);
        } else if (response.data && response.data.status === "error") {
          setValidationMessages(response.data.errors);
        }
      });
  };

  if (shouldRedirect) {
    return (
      <Redirect to="/" />
    );
  }

  const getValidationMessage = (key: string): string|null => {
    if (!validationMessages[key] || validationMessages[key].length === 0) {
      return null;
    }

    return validationMessages[key];
  };

  const handleEnter = (e: React.KeyboardEvent) => {
    if (e.charCode === 13) {
      submitForm();
    }
  };

  return (
    <Pane marginTop="15%" display="flex" alignItems="center" flexDirection="column">
      <Heading size={800}>Authenticate</Heading>
      <Pane
        elevation={3}
        backgroundColor="white"
        width={300}
        height={220}
        margin={24}
        display="flex"
        justifyContent="center"
        alignItems="center"
        flexDirection="column"
      >
        <TextInputField
          marginTop={15}
          marginBottom={20}
          height={35}
          label=""
          name="email"
          placeholder="Email"
          onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
            if (!e) {
              return;
            }
            const target = e.target as HTMLInputElement;
            if (target) {
              setEmail(target.value)
            }
          }}
          isInvalid={getValidationMessage("email") !== null}
          validationMessage={getValidationMessage("email")}
          onKeyPress={handleEnter}
        />
        <TextInputField
          marginTop={10}
          marginBottom={30}
          height={35}
          label=""
          name="password"
          type="password"
          placeholder="Password"
          onChange={ (e: React.ChangeEvent<HTMLInputElement>) => {
            if (!e) {
              return;
            }
            const target = e.target as HTMLInputElement;
            if (target) {
              setPassword(target.value)
            }
          }}
          isInvalid={getValidationMessage("password") !== null}
          validationMessage={getValidationMessage("password")}
          onKeyPress={handleEnter}
        />
        <Button
          onClick={ () => submitForm() }
        >
          Log in
        </Button>
      </Pane>
    </Pane>
  );
}

export default LogInForm;
