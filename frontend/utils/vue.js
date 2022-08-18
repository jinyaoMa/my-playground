// Usage
// export const ElMenu = withInstall(Menu, {
//   MenuItem,
//   MenuItemGroup,
//   SubMenu,
// })
export const withInstall = (main, extra) => {
  main.install = (app) => {
    for (const comp of [main, ...Object.values(extra ?? {})]) {
      app.component(comp.name, comp)
    }
  }

  if (extra) {
    for (const [key, comp] of Object.entries(extra)) {
      main[key] = comp
    }
  }
  return main
}

// Usage
// export const ElMenuItem = withNoopInstall(MenuItem)
// export const ElMenuItemGroup = withNoopInstall(MenuItemGroup)
// export const ElSubMenu = withNoopInstall(SubMenu)
export const withNoopInstall = (component) => {
  component.install = {}
  return component
}
