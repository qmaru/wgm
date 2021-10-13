import { useEffect, useCallback, useState } from 'react'
import Table from '@mui/material/Table'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import TableHead from '@mui/material/TableHead'
import TableRow from '@mui/material/TableRow'
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import Typography from '@mui/material/Typography'
import ButtonGroup from '@mui/material/ButtonGroup'
import Container from '@mui/material/Container'
import Paper from '@mui/material/Paper'
import Tooltip from '@mui/material/Tooltip'
import InputLabel from '@mui/material/InputLabel'
import MenuItem from '@mui/material/MenuItem'
import FormControl from '@mui/material/FormControl'
import Select, { SelectChangeEvent } from '@mui/material/Select'
import Checkbox from '@mui/material/Checkbox'
import FormGroup from '@mui/material/FormGroup'
import FormControlLabel from '@mui/material/FormControlLabel'
import Grid from '@mui/material/Grid'
import OutlinedInput from '@mui/material/OutlinedInput'
import Chip from '@mui/material/Chip'

import { DataWrapper, ConfirmWrapper } from './servers'
import { useSnackbar } from 'notistack'
import { DefaultMsgOption } from '../App'

export default function Clients() {
  const { enqueueSnackbar } = useSnackbar()
  const [isEdit, setIsEdit] = useState<boolean>(false)
  const [serverData, setServerData] = useState<any>([])
  const [clientData, setClientData] = useState<any>([])
  const [ruleData, setRuleData] = useState<any>([])

  const [dataTitle, setDataTitle] = useState<string>("")
  const [dataOpen, setDataOpen] = useState<boolean>(false)
  const [conOpen, setConOpen] = useState<boolean>(false)

  const [serverID, setServerID] = useState<string>("")
  const [clientID, setClientID] = useState<number>(0)
  const [clientName, setClientName] = useState<string>("")
  const [clientIP, setClientIP] = useState<string>("")
  const [clientExtra, setClientExtra] = useState<boolean>(false)
  const [clientCenter, setClientCenter] = useState<boolean>(false)
  const [clientAccess, setClientAccess] = useState<boolean>(false)
  const [clientKeep, setClientKeep] = useState<number>(25)

  const [routers, setRouters] = useState<string[]>([])
  const routersChange = (event: SelectChangeEvent<typeof routers>) => {
    const { target: { value } } = event
    setRouters(typeof value === 'string' ? value.split(',') : value)
  }

  const dataClean = () => {
    setDataOpen(false)
    setConOpen(false)
    setServerID("")
    setClientID(0)
    setClientName("")
    setClientIP("")
    setClientExtra(false)
    setClientCenter(false)
    setClientAccess(false)
    setClientKeep(25)
    setRouters([])
  }

  const dataClose = () => {
    setDataOpen(false)
  }

  const conClose = () => {
    setConOpen(false)
  }

  const serverIDChange = (event: SelectChangeEvent) => {
    setServerID(event.target.value)
  }

  const clientNameChange = (event: any) => {
    setClientName(event.target.value)
  }

  const clientIPChange = (event: any) => {
    setClientIP(event.target.value)
  }

  const clientExtraChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setClientExtra(event.target.checked)
  }

  const clientCenterChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setClientCenter(event.target.checked)
  }

  const clientAccessChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setClientAccess(event.target.checked)
  }

  const clientKeepChange = (event: any) => {
    setClientKeep(event.target.value)
  }

  const ClientAddOn = () => {
    setDataTitle("添加节点")
    setIsEdit(false)
    setDataOpen(true)
  }

  const ClientAddConfirm = () => {
    let client_data: any = {
      "username": clientName,
      "ip": clientIP,
      "keepalive": clientKeep,
    }

    if (clientAccess) {
      client_data["is_access"] = 1
    } else {
      client_data["is_access"] = 0
    }

    if (clientExtra) {
      client_data["is_extra"] = 1
    } else {
      client_data["is_extra"] = 0
    }

    if (clientCenter) {
      client_data["is_server"] = 1
    } else {
      client_data["is_server"] = 0
    }

    let url: string = `http://127.0.0.1:8373/api/v1/node/${serverID}`
    fetch(url, {
      method: "POST",
      body: JSON.stringify(client_data)
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("添加成功", DefaultMsgOption)
          dataClean()
          LoadClients()
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
        setDataOpen(false)
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const ClientKeyUpdateOn = (server_id: string, client_id: number, client_name: string) => {
    setConOpen(true)
    setServerID(server_id)
    setClientID(client_id)
    setClientName(client_name)
  }

  const ClientKeyUpdateConfirm = () => {
    const url: string = `http://127.0.0.1:8373/api/v1/node/${serverID}/${clientID}/key`
    fetch(url, {
      method: "PUT",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("修改密钥成功", DefaultMsgOption)
          LoadClients()
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
        setConOpen(false)
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const ClientEditOn = (server_id: string, client_id: number) => {
    setDataTitle("增加路由")
    setRouters([])
    setIsEdit(true)
    setDataOpen(true)
    setClientID(client_id)
    setServerID(server_id)
  }

  const ClientEditConfirm = () => {
    let ids: number[] = []
    for (let i in ruleData) {
      let rData: any = ruleData[i]
      routers.forEach(router => {
        if (rData.ip === router) {
          ids.push(rData.id)
        }
      })
    }

    if (ids.length === 0) {
      setDataOpen(false)
      return false
    }

    const idGroup: string = ids.join(",")
    const url: string = `http://127.0.0.1:8373/api/v1/data/userrule/${clientID}?rules=${idGroup}`
    fetch(url, {
      method: "PUT",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("添加路由", DefaultMsgOption)
        } else {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar(response.message, DefaultMsgOption)
        }
        setDataOpen(false)
      }).catch(
        () => {
          DefaultMsgOption.variant = "error"
          enqueueSnackbar("服务错误", DefaultMsgOption)
        }
      )
  }

  const ClientDel = (server_id: number, client_id: number) => {
    const url: string = `http://127.0.0.1:8373/api/v1/node/${server_id}/${client_id}`
    fetch(url, {
      method: "DELETE",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          let lastClientData: any = []
          for (let i in clientData) {
            let cData: any = clientData[i]
            if (cData["id"] !== client_id) {
              lastClientData.push(cData)
            }
          }
          DefaultMsgOption.variant = "success"
          enqueueSnackbar("删除成功", DefaultMsgOption)
          setClientData(lastClientData)
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

  const LoadClients = useCallback(() => {
    const url: string = `http://127.0.0.1:8373/api/v1/data/nodelist`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          const data: any = response.data
          setClientData(data)
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

  const LoadRules = useCallback(() => {
    const url: string = `http://127.0.0.1:8373/api/v1/data/rulelist`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        const status: number = response.status
        if (status === 1) {
          const data: any = response.data
          setRuleData(data)
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
    LoadClients()
    LoadRules()
  }, [LoadServers, LoadClients, LoadRules])

  let dataContent: any = []
  if (isEdit) {
    dataContent.push(
      <Box sx={{ p: 2 }} key="router">
        <FormControl sx={{ minWidth: 225 }}>
          <InputLabel>规则</InputLabel>
          <Select
            multiple
            value={routers}
            onChange={routersChange}
            input={<OutlinedInput label="Chip" />}
            renderValue={(selected) => (
              <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                {selected.map((value) => (
                  <Chip key={value} label={value} />
                ))}
              </Box>
            )}
            MenuProps={{
              PaperProps: {
                style: {
                  width: 225
                }
              }
            }}
          >
            {ruleData.map((data: any, index: number) => (
              <MenuItem
                key={"rule" + index}
                value={data.ip}
              >
                {data.ip}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </Box >
    )
  } else {
    dataContent.push(
      <Box
        key="clientInput"
        sx={{ flexGrow: 1, padding: 2 }}
      >
        <Grid
          container
          direction="column"
          justifyContent="center"
          alignItems="center"
          spacing={3}
        >
          <Grid item>
            <TextField
              required
              id="username"
              label="名称"
              value={clientName}
              onChange={clientNameChange}
            />
          </Grid>

          <Grid item>
            <FormControl sx={{ minWidth: 225 }}>
              <InputLabel>服务器</InputLabel>
              <Select
                value={serverID}
                label="服务器"
                onChange={serverIDChange}
              >
                {serverData.map((data: any, index: number) => {
                  return (
                    <MenuItem key={"server" + index} value={data.id}>{data.title} - {data.lan_ip}/{data.lan_netmask}</MenuItem>
                  )
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item>
            <TextField
              required
              id="ip"
              label="子网地址"
              value={clientIP}
              onChange={clientIPChange}
            />
          </Grid>

          <Grid item>
            <TextField
              required
              id="keepalive"
              label="心跳时间"
              value={clientKeep}
              onChange={clientKeepChange}
            />
          </Grid>

          <Grid item>
            <FormGroup sx={{ left: 36 }}>
              <FormControlLabel
                label="允许进入内网"
                control={
                  <Checkbox
                    checked={clientAccess}
                    onChange={clientAccessChange}
                  />
                }
              />
            </FormGroup>
          </Grid>

          <Grid item>
            <FormGroup sx={{ left: 36 }}>
              <FormControlLabel
                label="允许额外路由"
                control={
                  <Checkbox
                    checked={clientExtra}
                    onChange={clientExtraChange}
                  />
                }
              />
            </FormGroup>
          </Grid>

          <Grid item>
            <FormGroup sx={{ left: 36 }}>
              <FormControlLabel
                label="作为中心节点"
                control={
                  <Checkbox
                    checked={clientCenter}
                    onChange={clientCenterChange}
                  />
                }
              />
            </FormGroup>
          </Grid>
        </Grid>
      </Box>
    )
  }

  return (
    <Container sx={{ padding: 2 }}>
      <Button
        sx={{ margin: 2 }}
        variant="contained"
        onClick={() => ClientAddOn()}
      >
        添加节点
      </Button>

      {clientData.map((node: any, index: number) => {
        return (
          <Container sx={{ padding: 1 }}>
            <Typography variant="h5" sx={{ p: 1 }}>{node.server_title}</Typography>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell align="center">名称</TableCell>
                    <TableCell align="center">子网IP</TableCell>
                    <TableCell align="center">密钥</TableCell>
                    <TableCell align="center">默认路由</TableCell>
                    <TableCell align="center">额外路由</TableCell>
                    <TableCell align="center">心跳时间</TableCell>
                    <TableCell align="center">操作</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {node.users.map((data: any, index: number) => {
                    return (
                      <TableRow
                        key={"server-" + data.id}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                      >
                        <TableCell align="center">{data.username}</TableCell>
                        <TableCell align="center">{data.ip}</TableCell>
                        <TableCell align="center">
                          <ButtonGroup variant="text" color={data.is_server === 0 ? "primary" : "error"}>
                            <Tooltip disableFocusListener title={data.prikey}>
                              <Button>私钥</Button>
                            </Tooltip>
                            <Tooltip disableFocusListener title={data.pubkey}>
                              <Button>公钥</Button>
                            </Tooltip>
                          </ButtonGroup>
                        </TableCell>
                        <TableCell align="center">{data.default_rule}</TableCell>
                        <TableCell align="center">{data.is_extra === 0 ? "否" : "是"}</TableCell>
                        <TableCell align="center">{data.keepalive}</TableCell>

                        <TableCell align="center">
                          <ButtonGroup variant="text">
                            <Button
                              size='small'
                              disabled={data.is_extra === 0}
                              onClick={() => ClientEditOn(data.server_id, data.id)}
                            >路由</Button>
                            <Button
                              size='small'
                              onClick={() => ClientKeyUpdateOn(data.server_id, data.id, data.username)}
                            >更新</Button>
                            <Button
                              size='small'
                              color='error'
                              onClick={() => ClientDel(data.server_id, data.id)}
                            >删除</Button>
                          </ButtonGroup>
                        </TableCell>
                      </TableRow>
                    )
                  })}
                </TableBody>
              </Table>
            </TableContainer>
          </Container >
        )
      })}

      <DataWrapper
        dataOpen={dataOpen}
        dataClose={dataClose}
        dataTitle={dataTitle}
        dataContent={dataContent}
        dataSave={isEdit ? () => ClientEditConfirm() : () => ClientAddConfirm()}
      />

      <ConfirmWrapper
        open={conOpen}
        username={clientName}
        close={conClose}
        save={ClientKeyUpdateConfirm}
      />

    </Container >
  )
}