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
import VisibilityIcon from '@mui/icons-material/Visibility'
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff'

import { useSnackbar } from 'notistack'

export default function Users() {
  const { enqueueSnackbar } = useSnackbar()
  const [username, setUsername] = useState<string>("")
  const [userData, setUserData] = useState<any>([])
  const [showKey, setShowKey] = useState<boolean>(false)
  const [manualRender, setManualRender] = useState<boolean>(false)

  const UsernameChange = (event: any) => {
    setUsername(event.target.value)
  }

  const ShowSecret = () => {
    setShowKey(!showKey)
  }

  const UserAdd = () => {
    let body: any = {
      "username": username
    }
    const url = `${window.api}/user/add`
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
            "用户添加成功",
            window.messageDefault
          )
          setUsername("")
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
            "用户数据载入失败",
            window.messageDefault
          )
        }
      )
  }

  const UserList = useCallback(() => {
    const url = `${window.api}/user/list`
    fetch(url, {
      method: "GET",
    }).then(res => res.json())
      .then(response => {
        let status = response.status
        if (status === 1) {
          let data = response.data
          setUserData(data)
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
    UserList()
  }, [UserList, manualRender])

  return (
    <Container key={"Users-Main"}>
      <Container key={"Users-Control"}
        sx={{
          padding: 4,
          display: 'flex',
          justifyContent: 'center',
        }}
      >
        <Stack spacing={2}>
          <TextField
            label="用户名"
            variant="outlined"
            value={username}
            onChange={(event) => UsernameChange(event)}
          />
          <Button variant="contained" onClick={() => UserAdd()}>增加</Button>
        </Stack>
      </Container>

      <Container key={"Users-List"}>
        <Button
          variant="contained"
          onClick={() => ShowSecret()}
          startIcon={showKey ? <VisibilityOffIcon /> : <VisibilityIcon />}
        >
          {showKey ? "隐藏密钥" : "显示密钥"}
        </Button>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell sx={{ width: "4%" }} align="center">ID</TableCell>
              <TableCell sx={{ width: "10%" }} align="center">用户名</TableCell>
              <TableCell sx={{ width: "42%" }} align="center">私钥</TableCell>
              <TableCell sx={{ width: "42%" }} align="center">公钥</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {userData.map((data: any, index: number) => (
              <TableRow key={"userdata" + index}>
                <TableCell align="center">{data.id}</TableCell>
                <TableCell align="center">{data.username}</TableCell>
                {showKey && <TableCell align="center">{data.private_key}</TableCell>}
                {!showKey && <TableCell align="center">{"**********"}</TableCell>}
                {showKey && <TableCell align="center">{data.public_key}</TableCell>}
                {!showKey && <TableCell align="center">{"**********"}</TableCell>}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Container>
    </Container>
  )
}
