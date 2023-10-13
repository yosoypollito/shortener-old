import { ReactNode } from "react"

export interface ReactChildren{
  children: ReactNode;
}

export interface LinkState{
  step:number;
  link:string;
}

export interface IconProps{
  IconSVG:ReactNode;
  size?:string | number
}

export interface NavItemProps{
  to:string;
  label:string
}

export type Scope = string;

export interface Link {
  id:string;
  key:string;
  scope:Scope;
  dates:{
    creation:string
    modified:string
  }
}

export interface LinkResponse{
  link:Link
}

