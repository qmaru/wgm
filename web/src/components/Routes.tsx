import { useState, useEffect, useCallback } from 'react'

import Container from '@mui/material/Container'
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
import Stack from '@mui/material/Stack'
import Table from '@mui/material/Table'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'

import { useSnackbar } from 'notistack'

export default function Routes() {
  const { enqueueSnackbar } = useSnackbar()
  const [route, setRoute] = useState<string>("")
  const [routeData, setRouteData] = useState<any>([])
  const [manualRender, setManualRender] = useState<boolean>(false)

  const RouteChange = (event: any) => {
    setRoute(event.target.value)
  }

  const RouteAdd = () => {
    let body: any = {
      "cidr": route
    }
    const url = `${window.api}/route/add`
    fetch(url, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        'Content-Type': 'application/json',
      },
    }).then(res => res.json())
      .then(response => {
        let status = response.status
        if (status === 1) {
          setManualRender(!manualRender)
          window.messageDefault.variant = "success"
          enqueueSnackbar(
            "路由添加成功",
            window.messageDefault
          )
          setRoute("")
        } else {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "路由数据载入失败",
            window.messageDefault
          )
        }
      )
  }

  const RouteList = useCallback(() => {
    const url = `${window.api}/route/list`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        let status = response.status
        if (status === 1) {
          let data = response.data
          setRouteData(data)
        }
      })
      .catch(
        () => {
          window.messageDefault.variant = "error"
          enqueueSnackbar(
            "用户数据载入失败",
            window.messageDefault
          )
        }
      )
  }, [enqueueSnackbar])

  useEffect(() => {
    RouteList()
  }, [RouteList, manualRender])

  return (
    <Container key={"Routes-Main"}>
      <Container key={"Routes-Control"}
        sx={{
          padding: 4,
          display: 'flex',
          justifyContent: 'center',
        }}
      >
        <Stack spacing={2}>
          <TextField
            label="路由"
            variant="outlined"
            value={route}
            onChange={(event) => RouteChange(event)}
          />
          <Button variant="contained" onClick={() => RouteAdd()}>增加</Button>
        </Stack>
      </Container>

      <Container key={"Routes-List"} sx={{ width: 600 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell align="center">ID</TableCell>
              <TableCell align="center">CIDR</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {routeData.map((data: any, index: number) => (
              <TableRow key={"routedata" + index}>
                <TableCell align="center">{data.id}</TableCell>
                <TableCell align="center">{data.cidr}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Container>
    </Container>
  )
}
