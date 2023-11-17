import { useState, useMemo } from 'react'

import { createTheme, ThemeProvider } from '@mui/material/styles'
import themeColor from '@mui/material/colors/indigo'
import { SnackbarProvider } from 'notistack'
import useMediaQuery from '@mui/material/useMediaQuery'

import Tabs from '@mui/material/Tabs'
import Tab from '@mui/material/Tab'
import Container from '@mui/material/Container'

import Configs from './components/Configs'
import Peers from './components/Peers'
import Routes from './components/Routes'
import Users from './components/Users'

import "./global"

export default function App() {
  const prefersDarkMode: boolean = useMediaQuery('(prefers-color-scheme: dark)')
  const darkMode: boolean = prefersDarkMode

  var GlobalTheme = useMemo(
    () =>
      createTheme({
        palette: {
          mode: darkMode ? 'dark' : 'light',
          primary: {
            main: darkMode ? themeColor[600] : themeColor[700],
            contrastText: "#fff"
          },
          secondary: {
            main: darkMode ? themeColor[400] : themeColor[600],
            contrastText: "#fff"
          },
          contrastThreshold: 3,
          tonalOffset: 0.2,
        },
      }),
    [darkMode],
  )

  const [tabIndex, setTabIndex] = useState(0)
  const switchTab = (event: any, newValue: number) => {
    setTabIndex(newValue)
  }

  const tabList: any = [
    { label: "配置", comp: <Configs /> },
    { label: "节点", comp: <Peers /> },
    { label: "用户", comp: <Users /> },
    { label: "路由", comp: <Routes /> },
  ]

  return (
    <ThemeProvider theme={GlobalTheme}>
      <SnackbarProvider maxSnack={3} dense>
        <Container disableGutters maxWidth={false}
          sx={{
            position: 'relative',
            '::before': {
              content: `'v${window.version}'`,
              position: 'absolute',
              top: 14,
              left: 20,
              fontSize: 16,
              color: 'rgba(0, 0, 0, 0.7)',
              pointerEvents: 'none',
            },
          }}
        >
          <Container sx={{ borderBottom: 1, borderColor: 'divider' }} disableGutters maxWidth={false}>
            <Tabs value={tabIndex} onChange={switchTab} centered>
              {tabList.map((data: any, index: number) => {
                return (
                  <Tab
                    key={"tab-" + index}
                    label={data.label}
                  />
                )
              })}
            </Tabs>
          </Container >
          {tabList[tabIndex].comp}
        </Container>
      </SnackbarProvider>
    </ThemeProvider>
  )
}
