import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { authConfig } from "./auth.config";
import axios from "axios";
import { endPoint } from "@/constants";

const login = async (credentials) => {
  try {
    const response = await axios.post(endPoint + "/login", {
      email: credentials.email,
      password: credentials.password,
    });
    const { accessToken, _ } = response.data.data;

    const user = await fetch(endPoint + "/profile", {
      headers: {
        accept: "application/json",
        Authorization: `Bearer ${accessToken.token}`,
      },
    }).then((res) => res.json());
    user.token = accessToken;
    return user;
  } catch (e) {
    console.log(e.message);
    return null;
  }
};

export const {
  handlers: { GET, POST },
  auth,
  signIn,
  signOut,
} = NextAuth({
  ...authConfig,
  providers: [
    CredentialsProvider({
      async authorize(credentials, _) {
        try {
          const user = await login(credentials);
          return user;
        } catch (err) {
          return null;
        }
      },
    }),
  ],
  callbacks: {
    ...authConfig.callbacks,
  },
});
