import React from "react";
import { User } from "../model/user";

export const UserContext = React.createContext<User | undefined>(undefined);