import { deflate } from "zlib";

export interface MenuItem {
  title: string;
  danger: boolean;
  action: string;
}