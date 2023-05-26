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
import Tooltip from '@mui/material/Tooltip'

import { useSnackbar } from 'notistack'

export default function Users() {
  const { enqueueSnackbar } = useSnackbar()
  const [manualRender, setManualRender] = useState<boolean>(false)

  const [userData, setUserData] = useState<any>([])

  const [userAddUsername, setUserAddUsername] = useState<string>("")

  const [userUpdateOpen, setUserUpdateOepn] = useState<boolean>(false)
  const [userUpdateID, setUserUpdateID] = useState<number>(0)
  const [userUpdateUsername, setUserUpdateUsername] = useState<string>("")

  const [userDeleteOpen, setUserDeleteOepn] = useState<boolean>(false)
  const [userDeleteID, setUserDeleteID] = useState<number>(0)


  const UserAddUsernameChange = (event: any) => {
    setUserAddUsername(event.target.value)
  }

  const UserUpdateOpen = (user_data: any) => {
    setUserUpdateOepn(true)
    setUserUpdateID(user_data.id)
    setUserUpdateUsername(user_data.username)
  }

  const UserUpdateClose = () => {
    setUserUpdateOepn(false)
  }

  const UserDeleteOpen = (user_data: any) => {
    setUserDeleteOepn(true)
    setUserDeleteID(user_data.id)
  }

  const UserDeleteClose = () => {
    setUserDeleteOepn(false)
  }

  const UserUpdateUsernameChange = (event: any) => {
    setUserUpdateUsername(event.target.value)
  }

  const UserAdd = () => {
    let body: any = {
      "username": userAddUsername
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
            response.message,
            window.messageDefault
          )
          setUserAddUsername("")
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
            "用户接口请求失败",
            window.messageDefault
          )
        }
      )
  }

  const UserUpdate = () => {
    let body: any = {
      "username": userUpdateUsername
    }
    const url = `${window.api}/user/update/` + userUpdateID
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
          setUserUpdateOepn(false)
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
            "用户接口请求失败",
            window.messageDefault
          )
        }
      )
  }

  const UserDelete = () => {
    const url = `${window.api}/user/delete/` + userDeleteID
    fetch(url, {
      method: "POST",
    }).then(res => res.json())
      .then(response => {
        let status = response.status
        if (status === 1) {
          setManualRender(!manualRender)
          setUserDeleteOepn(false)
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
            "用户接口请求失败",
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
            "用户接口请求失败",
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
            value={userAddUsername}
            onChange={(event) => UserAddUsernameChange(event)}
          />
          <Button variant="contained" onClick={() => UserAdd()}>提交</Button>
        </Stack>
      </Container>

      <Container key={"Users-List"}>
        <Stack
          spacing={{ xs: 2, sm: 2 }}
          direction="row"
          justifyContent="flex-start"
          useFlexGap
          flexWrap="wrap"
        >
          {userData.map((data: any, index: number) => (
            <Card key={"user" + index} sx={{ minWidth: 200 }}>
              <CardContent sx={{ textAlign: "center" }}>
                <Typography variant="subtitle2" gutterBottom>
                  {data.username}
                </Typography>
                <Tooltip title={data.private_key}>
                  <Button>私钥</Button>
                </Tooltip>
                <Tooltip title={data.public_key}>
                  <Button>公钥</Button>
                </Tooltip>
              </CardContent>
              <CardActions>
                <Button onClick={() => UserUpdateOpen(data)}>修改</Button>
                <Button onClick={() => UserDeleteOpen(data)} color="error" >删除</Button>
              </CardActions>
            </Card>
          ))}
        </Stack>
      </Container>

      <Dialog open={userUpdateOpen} onClose={UserUpdateClose}>
        <DialogTitle>修改用户</DialogTitle>
        <DialogContent>
          <Box sx={{ padding: 1 }}>
            <TextField
              label="用户名"
              variant="outlined"
              value={userUpdateUsername}
              onChange={(event) => UserUpdateUsernameChange(event)}
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={UserUpdateClose}>取消</Button>
          <Button onClick={() => UserUpdate()}>提交</Button>
        </DialogActions>
      </Dialog>


      <Dialog open={userDeleteOpen} onClose={UserDeleteClose}>
        <DialogTitle>确认删除用户</DialogTitle>
        <DialogActions>
          <Button onClick={UserDeleteClose}>取消</Button>
          <Button onClick={() => UserDelete()}>提交</Button>
        </DialogActions>
      </Dialog>


    </Container>
  )
}