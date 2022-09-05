type ExternalApps = {
  key: string;
  title: Record<string, string>;
  link: string;
  icon: string;
  vitepress?: boolean;
}[];

const eas: ExternalApps = [
  {
    key: "Docs",
    title: {
      zh: "文档",
      en: "Docs",
    },
    link: "/docs/{lang}/index.html",
    icon: "docs",
    vitepress: true,
  },
];

// @ts-ignore
// export default process.env.NODE_ENV != "production" ? eas : [];
// export default process.env.NODE_ENV === "production" ? eas : [];
export default eas;
