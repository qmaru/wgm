import { useState, useMemo } from 'react'

import { createTheme, ThemeProvider } from '@mui/material/styles'
import { blue } from '@mui/material/colors'
import { SnackbarProvider } from 'notistack'

import Tabs from '@mui/material/Tabs'
import Tab from '@mui/material/Tab'
import Container from '@mui/material/Container'

import Peers from './components/Peers'
import Routes from './components/Routes'
import Users from './components/Users'

import "./global"

export default function App() {
  var GlobalTheme = useMemo(
    () =>
      createTheme({
        palette: {
          primary: {
            main: blue[800],
          },
        },
      }),
    [],
  )

  const [tabIndex, setTabIndex] = useState(0)
  const switchTab = (event: any, newValue: number) => {
    setTabIndex(newValue)
  }

  const tabList: any = [
    { label: "节点", comp: <Peers /> },
    { label: "用户", comp: <Users /> },
    { label: "路由", comp: <Routes /> },
  ]

  return (
    <ThemeProvider theme={GlobalTheme}>
      <SnackbarProvider maxSnack={3} dense>
        <Container disableGutters>
          <Container sx={{ borderBottom: 1, borderColor: 'divider' }}>
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
