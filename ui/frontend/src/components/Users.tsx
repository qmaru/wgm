import { useState, useEffect, useCallback } from "react";

import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import Tooltip from "@mui/material/Tooltip";
import Chip from "@mui/material/Chip";
import { useSnackbar } from "notistack";

import {
  UserListAPI,
  UserAddAPI,
  UserUpdateAPI,
  UserDeleteAPI,
} from "../../wailsjs/go/backend/App";

import { MyCard } from "./common";

export default function Users() {
  const { enqueueSnackbar } = useSnackbar();
  const [manualRender, setManualRender] = useState<boolean>(false);

  const [userData, setUserData] = useState<any>([]);

  const [userAddUsername, setUserAddUsername] = useState<string>("");

  const [userUpdateOpen, setUserUpdateOepn] = useState<boolean>(false);
  const [userUpdateID, setUserUpdateID] = useState<number>(0);
  const [userUpdateUsername, setUserUpdateUsername] = useState<string>("");

  const [userDeleteOpen, setUserDeleteOepn] = useState<boolean>(false);
  const [userDeleteID, setUserDeleteID] = useState<number>(0);

  const UserAddUsernameChange = (event: any) => {
    setUserAddUsername(event.target.value);
  };

  const UserUpdateOpen = (user_data: any) => {
    setUserUpdateOepn(true);
    setUserUpdateID(user_data.id);
    setUserUpdateUsername(user_data.username);
  };

  const UserUpdateClose = () => {
    setUserUpdateOepn(false);
  };

  const UserDeleteOpen = (user_data: any) => {
    setUserDeleteOepn(true);
    setUserDeleteID(user_data.id);
  };

  const UserDeleteClose = () => {
    setUserDeleteOepn(false);
  };

  const UserUpdateUsernameChange = (event: any) => {
    setUserUpdateUsername(event.target.value);
  };

  const UserAdd = () => {
    let body: any = {
      username: userAddUsername,
    };
    UserAddAPI(body)
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          setManualRender(!manualRender);
          window.messageDefault.variant = "success";
          enqueueSnackbar(response.message, window.messageDefault);
          setUserAddUsername("");
        } else {
          window.messageDefault.variant = "error";
          enqueueSnackbar(response.message, window.messageDefault);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("用户接口请求失败", window.messageDefault);
      });
  };

  const UserUpdate = () => {
    let body: any = {
      username: userUpdateUsername,
    };

    UserUpdateAPI(String(userUpdateID), body)
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          setManualRender(!manualRender);
          setUserUpdateOepn(false);
          window.messageDefault.variant = "success";
          enqueueSnackbar(response.message, window.messageDefault);
        } else {
          window.messageDefault.variant = "error";
          enqueueSnackbar(response.message, window.messageDefault);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("用户接口请求失败", window.messageDefault);
      });
  };

  const UserDelete = () => {
    UserDeleteAPI(String(userDeleteID))
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          setManualRender(!manualRender);
          setUserDeleteOepn(false);
          window.messageDefault.variant = "success";
          enqueueSnackbar(response.message, window.messageDefault);
        } else {
          window.messageDefault.variant = "error";
          enqueueSnackbar(response.message, window.messageDefault);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("用户接口请求失败", window.messageDefault);
      });
  };

  const UserList = useCallback(() => {
    UserListAPI()
      .then((response) => {
        let status = response.status;
        if (status === 1) {
          let data = response.data;
          setUserData(data);
        }
      })
      .catch(() => {
        window.messageDefault.variant = "error";
        enqueueSnackbar("用户接口请求失败", window.messageDefault);
      });
  }, [enqueueSnackbar]);

  useEffect(() => {
    UserList();
  }, [UserList, manualRender]);

  return (
    <Container key={"Users-Main"} maxWidth={false}>
      <Container
        key={"Users-Control"}
        disableGutters
        maxWidth={false}
        sx={{
          padding: 4,
          display: "flex",
          justifyContent: "center",
        }}
      >
        <Stack spacing={2}>
          <TextField
            label="用户名"
            variant="outlined"
            value={userAddUsername}
            onChange={(event) => UserAddUsernameChange(event)}
          />
          <Button variant="contained" onClick={() => UserAdd()}>
            提交
          </Button>
        </Stack>
      </Container>

      <Container key={"Users-List"} disableGutters sx={{ paddingBottom: 4, maxWidth: 800 }}>
        <Stack
          spacing={{ xs: 2, sm: 2 }}
          direction="row"
          justifyContent="flex-start"
          useFlexGap
          flexWrap="wrap"
        >
          {userData.map((data: any, index: number) => (
            <MyCard
              key={"user" + index}
              content={
                <>
                  <Typography variant="body1" gutterBottom>
                    {data.username}
                  </Typography>
                  <Stack spacing={2} justifyContent="center" alignItems="center" sx={{ p: 2 }}>
                    <Tooltip title={data.private_key} placement="top">
                      <Chip
                        clickable
                        sx={{ borderRadius: 2, width: 100 }}
                        variant="outlined"
                        label="Private"
                        color="info"
                      />
                    </Tooltip>
                    <Tooltip title={data.public_key} placement="top">
                      <Chip
                        clickable
                        sx={{ borderRadius: 2, width: 100 }}
                        variant="outlined"
                        label="Public"
                        color="success"
                      />
                    </Tooltip>
                  </Stack>
                </>
              }
              contentStyle={{ pb: 4, minWidth: 150 }}
              onEdit={() => UserUpdateOpen(data)}
              onDelete={() => UserDeleteOpen(data)}
            />
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
  );
}
