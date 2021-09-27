# ce-ssh

Two parts:

- udp-client -> udp->tcp proxy -> ssh client/tunnel -> sshd (on a VM)
- sshd (VM) -> tcp->udp proxy -> udp_service

The server (sshd) side of things is run in a set of Code Engine jobs.
One for sshd (to simulate a VM), and one for the UDP listening service.
The VM will have 2 processes running:
- sshd
- netcat

Netcat will proxy tcp connections to udp, which then gets forwarded on
to the udp_service.

The sshd process will just wait for incoming ssh connections.

The ssh client/tunnel will do two things:
- create an ssh connection to the fake VM
- create a udp->tdp proxy so our udp-client can send udp messages to the
  udp->tcp proxy, which then forwards on to the ssh (tcp) tunnel

The entire ssh key sharing thing you'll see in the vm setup is just for
demo purposes. Do not do share ssh keys this way in a real env.

# Do it!

Create a Code Engine project and make sure it's selected:
```
$ ibmcloud ce project create -n demo
```

Then create the fake VM:
```
$ cd vm
$ ./run
```

Now run the client app that will send UDP messages:
```
$ cd ../app
$ ./run
```

If everything worked then you should see the output of the client app like this:
```
2021/09/28 00:38:21 Connecting to: 172.30.165.203:1234
Reply: yes I got:From client message #0
Reply: yes I got:From client message #1
Reply: yes I got:From client message #2
Reply: yes I got:From client message #3
Reply: yes I got:From client message #4
Reply: yes I got:From client message #5
Reply: yes I got:From client message #6
Reply: yes I got:From client message #7
Reply: yes I got:From client message #8
Reply: yes I got:From client message #9
Job run completed successfully.
```

# Clean up

```
$ vm/run clean
$ app/run clean
```
