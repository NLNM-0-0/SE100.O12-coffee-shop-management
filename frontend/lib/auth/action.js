"use server"

import { auth, signIn, signOut } from "./auth";

export const login = async (prevState, formData) => {
    const { email, password } = Object.fromEntries(formData);

    try {
        await signIn("credentials", { email, password });
    } catch (err) {
        console.log(err);

        if (err.message.includes("CredentialsSignin")) {
            return { error: "Invalid username or password" };
        }
        throw err;
    }
    console.log("Log in successfully -----------------------")
};

export const logOut = async () => {
    await signOut();
};

export const getApiKey = async () => {
    const session = await auth()
    return session?.user?.token
}

export const getUser = async () => {
    const session = await auth()
    return session?.user
}