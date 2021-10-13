import { useEffect, useCallback, useState } from 'react'
import Container from '@mui/material/Container'
import Table from '@mui/material/Table'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'
import Paper from '@mui/material/Paper'
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import ButtonGroup from '@mui/material/ButtonGroup'
import Dialog from '@mui/material/Dialog'
import DialogActions from '@mui/material/DialogActions'
import DialogContent from '@mui/material/DialogContent'
import DialogTitle from '@mui/material/DialogTitle'

import { useSnackbar } from 'notistack'
import { DefaultMsgOption } from '../App'

export const DataWrapper = (props: any) => {
  return (
    <Dialog
      open={props.dataOpen}
      onClose={props.dataClose}
    >
      <DialogTitle>{props.dataTitle}</DialogTitle>
      <DialogContent>
        {props.dataContent}
      </DialogContent>
      <DialogActions>
        <Button onClick={props.dataClose}>取消</Button>
        <Button onClick={props.dataSave}>提交</Button>
      </DialogActions>
    </Dialog>
  )
}

export const ConfirmWrapper = (props: any) => {
  return (
    <Dialog
      open={props.open}
      onClose={props.close}
    >
      <DialogTitle>
        确定更新[{props.username}]的密钥
      </DialogTitle>
      <DialogActions>
        <Button onClick={props.close}>取消</Button>
        <Button onClick={props.save} autoFocus>确定</Button>
      </DialogActions>
    </Dialog>
  )
}

export default function Servers() {
  const { enqueueSnackbar } = useSnackbar()
  const [isEdit, setIsEdit] = useState<boolean>(false)
  const [serverData, setServerData] = useState<any>([])
  const [serverID, setServerID] = useState<number>(0)
  const [serverTitle, setServerTitle] = useState<string>("")
  const [serverAddress, setServerAddress] = useState<string>("")
  const [serverPort, setServerPort] = useState<number>(443)
  const [serverLanIP, setServerLanIP] = useState<string>("")
  const [serverLanNetmask, setServerLanNetmask] = useState<string>("")
  const [serverMTU, setServerMTU] = useState<number>(0)
  const [serverDNS, setServerDNS] = useState<string>("")

  const [dataTitle, setDataTitle] = useState<string>("")
  const [dataOpen, setDataOpen] = useState<boolean>(false)

  const serverTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setServerTitle(event.target.value)
  }

  const serverAddressChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setServerAddress(event.target.value)
  }

  const serverPortChange = (event: any) => {
    setServerPort(event.target.value)
  }

  const serverLanIPChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setServerLanIP(event.target.value)
  }

  const serverLanNetmaskChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setServerLanNetmask(event.target.value)
  }

  const serverMTUChange = (event: any) => {
    setServerMTU(event.target.value)
  }

  const serverDNSChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setServerDNS(event.target.value)
  }

  const dataClose = () => {
    dataClean()
    setDataOpen(false)
  }

  const dataClean = () => {
    setServerTitle("")
    setServerAddress("")
    setServerPort(443)
    setServerLanIP("")
    setServerLanNetmask("")
    setServerMTU(0)
    setServerDNS("")
  }

  const ServerAddOn = () => {
    setDataTitle("添加服务器")
    setDataOpen(true)
  }

  const ServerAddConfirm = () => {
    let server_data: any = {
      "title": serverTitle,
      "address": serverAddress,
      "port": serverPort,
      "lan_ip": serverLanIP,
      "lan_netmask": serverLanNetmask
    }

    if (serverMTU !== 0) {
      server_data["mtu"] = serverMTU
    }

    if (serverDNS !== "") {
      server_data["dns"] = serverDNS
    }

    let url: string = `http://127.0.0.1:8373/api/v1/server`
    fetch(url, {
      method: "POST",
      body: JSON.stringify(server_data)
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("添加成功", DefaultMsgOption)
          setDataOpen(false)
          dataClean()
          LoadServers()
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const ServerEditOn = (data: any) => {
    setDataTitle("编辑服务器")
    setServerID(data.id)
    setServerTitle(data.title)
    setServerAddress(data.address)
    setServerPort(data.port)
    setServerLanIP(data.lan_ip)
    setServerLanNetmask(data.lan_netmask)
    setServerMTU(data.mtu)
    setServerDNS(data.dns)
    setIsEdit(true)
    setDataOpen(true)
  }

  const ServerEditConfirm = () => {
    setIsEdit(false)

    let server_data: any = {
      "title": serverTitle,
      "address": serverAddress,
      "port": serverPort,
      "lan_ip": serverLanIP,
      "lan_netmask": serverLanNetmask
    }

    if (serverMTU !== 0) {
      server_data["mtu"] = serverMTU
    }

    if (serverDNS !== "") {
      server_data["dns"] = serverDNS
    }

    let url: string = `http://127.0.0.1:8373/api/v1/server/${serverID}`
    fetch(url, {
      method: "PUT",
      body: JSON.stringify(server_data)
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("修改成功", DefaultMsgOption)
          setDataOpen(false)
          dataClean()
          LoadServers()
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const ServerDel = (server_id: number) => {
    let url: string = `http://127.0.0.1:8373/api/v1/server/${server_id}`
    fetch(url, {
      method: "DELETE",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          let lastServerData: any = []
          for (let i in serverData) {
            let sData: any = serverData[i]
            if (sData["id"] !== server_id) {
              lastServerData.push(sData)
            }
          }
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("删除成功", DefaultMsgOption)
          setServerData(lastServerData)
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const LoadServers = useCallback(() => {
    const url: string = `http://127.0.0.1:8373/api/v1/data/serverlist`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          const data: any = response.data
          setServerData(data)
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }, [enqueueSnackbar])

  useEffect(() => {
    LoadServers()
  }, [LoadServers])

  let dataContent: any = []
  dataContent.push(
    <Box
      key="serverInput"
      component="form"
      sx={{
        '& .MuiTextField-root': { m: 2, width: '25ch' },
      }}
      autoComplete="off"
    >
      <div>
        <TextField
          required
          id="title"
          label="名称"
          value={serverTitle}
          onChange={serverTitleChange}
        />
        <TextField
          required
          id="address"
          label="服务器地址"
          value={serverAddress}
          onChange={serverAddressChange}
        />
        <TextField
          required
          id="port"
          label="服务器端口"
          type="number"
          value={serverPort}
          onChange={serverPortChange}
        />
        <TextField
          required
          id="lan_ip"
          label="子网网关"
          value={serverLanIP}
          onChange={serverLanIPChange}
        />
        <TextField
          required
          id="lan_netmask"
          label="子网掩码"
          value={serverLanNetmask}
          onChange={serverLanNetmaskChange}
        />
        <TextField
          required
          id="mtu"
          label="MTU"
          type="number"
          value={serverMTU}
          onChange={serverMTUChange}
        />
        <TextField
          required
          id="dns"
          label="DNS"
          value={serverDNS}
          onChange={serverDNSChange}
        />
      </div>
    </Box>
  )

  return (
    <Container sx={{ padding: 2 }}>
      <Button
        sx={{ margin: 2 }}
        variant="contained"
        onClick={() => ServerAddOn()}
      >
        添加服务器
      </Button>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell align="center">ID</TableCell>
              <TableCell align="center">名称</TableCell>
              <TableCell align="center">地址</TableCell>
              <TableCell align="center">端口</TableCell>
              <TableCell align="center">子网网关</TableCell>
              <TableCell align="center">子网掩码</TableCell>
              <TableCell align="center">MTU</TableCell>
              <TableCell align="center">DNS</TableCell>
              <TableCell align="center">操作</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {serverData.map((data: any, index: number) => {
              return (
                <TableRow
                  key={"server-" + data.id}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  <TableCell align="center">{data.id}</TableCell>
                  <TableCell align="center">{data.title}</TableCell>
                  <TableCell align="center">{data.address}</TableCell>
                  <TableCell align="center">{data.port}</TableCell>
                  <TableCell align="center">{data.lan_ip}</TableCell>
                  <TableCell align="center">{data.lan_netmask}</TableCell>
                  <TableCell align="center">{data.mtu}</TableCell>
                  <TableCell align="center">{data.dns}</TableCell>

                  <TableCell align="center">
                    <ButtonGroup variant="contained">
                      <Button
                        size='small'
                        onClick={() => { ServerEditOn(data) }}
                      >
                        编辑
                      </Button>
                      <Button
                        size='small'
                        variant="contained"
                        color='error'
                        onClick={() => ServerDel(data.id)}
                      >
                        删除
                      </Button>
                    </ButtonGroup>
                  </TableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>

      <DataWrapper
        dataOpen={dataOpen}
        dataClose={dataClose}
        dataTitle={dataTitle}
        dataContent={dataContent}
        dataSave={isEdit ? ServerEditConfirm : ServerAddConfirm}
      />
    </Container>
  )
}