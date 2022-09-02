type ExternalApps = {
  key: string;
  title: Record<string, string>;
  link: string;
}[];

const eas: ExternalApps = [
  {
    key: "Docs",
    title: {
      zh: "文档",
      en: "Docs",
    },
    link: "/docs/index.html",
  },
];

// @ts-ignore
// export default process.env.NODE_ENV != "production" ? eas : [];
export default process.env.NODE_ENV === "production" ? eas : [];
