type ExternalApps = {
  key: string;
  title: Record<string, string>;
  link: string;
  vitepress?: boolean;
}[];

const eas: ExternalApps = [
  {
    key: "Docs",
    title: {
      zh: "文档",
      en: "Docs",
    },
    link: "/docs/index.html",
    vitepress: true,
  },
];

// @ts-ignore
// export default process.env.NODE_ENV != "production" ? eas : [];
// export default process.env.NODE_ENV === "production" ? eas : [];
export default eas;
