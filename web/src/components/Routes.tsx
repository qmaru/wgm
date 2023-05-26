import { useState, useEffect, useCallback } from 'react'

import Container from '@mui/material/Container'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
import Stack from '@mui/material/Stack'
import Typography from '@mui/material/Typography'
import Card from '@mui/material/Card'
import CardActions from '@mui/material/CardActions'
import CardContent from '@mui/material/CardContent'
import Dialog from '@mui/material/Dialog'
import DialogActions from '@mui/material/DialogActions'
import DialogContent from '@mui/material/DialogContent'
import DialogTitle from '@mui/material/DialogTitle'

import { useSnackbar } from 'notistack'

export default function Routes() {
  const { enqueueSnackbar } = useSnackbar()
  const [manualRender, setManualRender] = useState<boolean>(false)

  const [routeData, setRouteData] = useState<any>([])

  const [routeAddCIDR, setRouteAddCIDR] = useState<string>("")

  const [routeUpdateOpen, setRouteUpdateOepn] = useState<boolean>(false)
  const [routeUpdateID, setRouteUpdateID] = useState<number>(0)
  const [routeUpdateCIDR, setRouteUpdateCIDR] = useState<string>("")

  const [routeDeleteOpen, setRouteDeleteOepn] = useState<boolean>(false)
  const [routeDeleteID, setRouteDeleteID] = useState<number>(0)

  const RouteAddChange = (event: any) => {
    setRouteAddCIDR(event.target.value)
  }

  const RouteUpdateChange = (event: any) => {
    setRouteUpdateCIDR(event.target.value)
  }

  const RouteUpdateOpen = (route_data: any) => {
    setRouteUpdateOepn(true)
    setRouteUpdateID(route_data.id)
    setRouteUpdateCIDR(route_data.cidr)
  }

  const RouteUpdateClose = () => {
    setRouteUpdateOepn(false)
  }

  const RouteDeleteOpen = (route_data: any) => {
    setRouteDeleteOepn(true)
    setRouteDeleteID(route_data.id)
  }

  const RouteDeleteClose = () => {
    setRouteDeleteOepn(false)
  }

  const RouteAdd = () => {
    let body: any = {
      "cidr": routeAddCIDR
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
            response.message,
            window.messageDefault
          )
          setRouteAddCIDR("")
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
            "路由接口请求失败",
            window.messageDefault
          )
        }
      )
  }

  const RouteUpdate = () => {
    let body: any = {
      "cidr": routeUpdateCIDR
    }
    const url = `${window.api}/route/update/` + routeUpdateID
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
          setRouteUpdateOepn(false)
          window.messageDefault.variant = "success"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
          setRouteAddCIDR("")
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
            "路由接口请求失败",
            window.messageDefault
          )
        }
      )
  }

  const RouteDelete = () => {
    const url = `${window.api}/route/delete/` + routeDeleteID
    fetch(url, {
      method: "POST",
    }).then(res => res.json())
      .then(response => {
        let status = response.status
        if (status === 1) {
          setManualRender(!manualRender)
          setRouteDeleteOepn(false)
          window.messageDefault.variant = "success"
          enqueueSnackbar(
            response.message,
            window.messageDefault
          )
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
            "路由接口请求失败",
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
            "路由接口请求失败",
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
            value={routeAddCIDR}
            onChange={(event) => RouteAddChange(event)}
          />
          <Button variant="contained" onClick={() => RouteAdd()}>提交</Button>
        </Stack>
      </Container>

      <Container key={"Routes-List"} sx={{ width: 600 }}>
        <Stack
          spacing={{ xs: 2, sm: 2 }}
          direction="row"
          justifyContent="flex-start"
          useFlexGap
          flexWrap="wrap"
        >
          {routeData.map((data: any, index: number) => (
            <Card key={"route" + index}>
              <CardContent>
                <Typography variant="subtitle2" gutterBottom>
                  {data.cidr}
                </Typography>
              </CardContent>
              <CardActions>
                <Button onClick={() => RouteUpdateOpen(data)}>修改</Button>
                <Button onClick={() => RouteDeleteOpen(data)} color="error" >删除</Button>
              </CardActions>
            </Card>
          ))}
        </Stack>
      </Container>

      <Dialog open={routeUpdateOpen} onClose={RouteUpdateClose}>
        <DialogTitle>修改路由</DialogTitle>
        <DialogContent>
          <Box sx={{ padding: 1 }}>
            <TextField
              label="路由"
              variant="outlined"
              value={routeUpdateCIDR}
              onChange={(event) => RouteUpdateChange(event)}
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={RouteUpdateClose}>取消</Button>
          <Button onClick={() => RouteUpdate()}>提交</Button>
        </DialogActions>
      </Dialog>


      <Dialog open={routeDeleteOpen} onClose={RouteDeleteClose}>
        <DialogTitle>确认删除路由</DialogTitle>
        <DialogActions>
          <Button onClick={RouteDeleteClose}>取消</Button>
          <Button onClick={() => RouteDelete()}>提交</Button>
        </DialogActions>
      </Dialog>

    </Container>
  )
}
