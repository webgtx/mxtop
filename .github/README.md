# mxtop
The Top Multiplexer on steroids

### Requirements
- Go
- Ansible
- Any unix-family OS

### Ansible infrastructure
```bash
.
├── inventory
│   └── hosts # Here you can configure your hosts list 
├── LICENSE
├── main.yml # There you should write your $remote_user_name, and select hosts groups
└── roles
    └── remotefetch
        └── tasks
            └── main.yml
```

#### Samples for playbooks configuration

**main.yml**
```yml
---
- hosts: all # You must edit this two arguments
  remote_user: dxv1d # and this one

  roles:
    - remotefetch
```

**inventory/hosts**
> As you noticed, this is plain ini configuration file. Use "[]" and write your IP  
```
[redhat]
192.168.1.99

[oracle]
192.168.1.117

[all]
192.168.1.99
192.168.1.117
```

#### How to run this playbooks?
```bash
ansible-playbook main.yml -i inventory/hosts
```
