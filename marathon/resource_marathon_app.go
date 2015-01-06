package marathon

import (
	"fmt"
	"github.com/Banno/go-marathon"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"time"
)

func resourceMarathonApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceMarathonAppCreate,
		Read:   resourceMarathonAppRead,
		Update: resourceMarathonAppUpdate,
		Delete: resourceMarathonAppDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{ // represents 'id' field
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"args": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backoff_seconds": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"backoff_factor": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"cmd": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"constraints": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"constraint": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"operation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"parameter": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"container": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"docker": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"network": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_mappings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: false,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port_mapping": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: false,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"container_port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"host_port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"service_port": &schema.Schema{
																Type:     schema.TypeInt,
																Optional: true,
															},
															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"volumes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"volume": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: false,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"container_path": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"host_path": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"mode": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"cpus": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"dependencies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disk": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"env": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"health_checks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"path": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"grace_period_seconds": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interval_seconds": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"port_index": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"timeout_seconds": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"max_consecutive_failures": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"command": &schema.Schema{
										Type:     schema.TypeMap,
										Optional: true,
									},
									// incomplete computed values here
								},
							},
						},
					},
				},
			},
			"instances": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeString, //should be float -_-
				Optional: true,
				ForceNew: false,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"upgrade_strategy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"minimum_health_capacity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"uris": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			// many other "computed" values haven't been added.
		},
	}
}

func resourceMarathonAppCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	appMutable := mutateResourceToAppMutable(d)

	app, err := c.AppCreate(appMutable)
	if err != nil {
		log.Println(err)
		return err
	}

	d.SetId(app.Id)

	// inspect the returned App stuff and set more computed values

	return resourceMarathonAppRead(d, meta)
}

func resourceMarathonAppRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	// client should throw error if id is nil
	app, _ := c.AppRead(d.Id())

	if app.Id == "" {
		d.SetId("")
	}

	// Add in computed values from App struct here.
	d.Set("version", app.Version)

	return nil
}

func resourceMarathonAppUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	appMutable := mutateResourceToAppMutable(d)

	// Spin until it's not locked by a deployment or things time out.
	stateConf := &resource.StateChangeConf{
		Pending:    []string{""},
		Target:     "updated",
		Refresh:    updateAppFunc(c, &appMutable),
		Timeout:    10 * time.Minute,
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Timed out, or something: %#v", err)
	}

	time.Sleep(5 * time.Second)

	return resourceMarathonAppRead(d, meta)
}

func resourceMarathonAppDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*marathon.Client)

	if err := c.AppDelete(d.Id()); err != nil {
		return err
	}

	return nil
}

func updateAppFunc(c *marathon.Client, appMutable *marathon.AppMutable) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		appUpdateResponse, err := c.AppUpdate(*appMutable, false)
		if err != nil {
			log.Printf("Update Error: %#v\n", err)
			return nil, "", err
		}

		return appUpdateResponse, "updated", nil
	}
}

func mutateResourceToAppMutable(d *schema.ResourceData) marathon.AppMutable {

	appMutable := marathon.AppMutable{}

	if v, ok := d.GetOk("name"); ok {
		appMutable.Id = v.(string)
	}

	if v, ok := d.GetOk("args.#"); ok {
		args := make([]string, v.(int))

		for i, _ := range args {
			args[i] = d.Get("args." + strconv.Itoa(i)).(string)
		}

		if len(args) != 0 {
			appMutable.Args = args
		}
	}

	if v, ok := d.GetOk("backoff_seconds"); ok {
		backoffSeconds, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.BackoffSeconds = backoffSeconds
	}

	if v, ok := d.GetOk("backoff_factor"); ok {
		backoffFactor, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.BackoffFactor = backoffFactor
	}

	if v, ok := d.GetOk("cmd"); ok {
		appMutable.Cmd = v.(string)
	}

	if v, ok := d.GetOk("constraints.0.constraint.#"); ok {
		constraints := make([][]string, v.(int))

		for i, _ := range constraints {
			cMap := d.Get(fmt.Sprintf("constraints.0.constraint.%d", i)).(map[string]interface{})

			if cMap["parameter"] == "" {
				constraints[i] = make([]string, 2)
				constraints[i][0] = cMap["attribute"].(string)
				constraints[i][1] = cMap["operation"].(string)
			} else {
				constraints[i] = make([]string, 3)
				constraints[i][0] = cMap["attribute"].(string)
				constraints[i][1] = cMap["operation"].(string)
				constraints[i][2] = cMap["parameter"].(string)
			}
		}

		appMutable.Constraints = constraints
	}

	if v, ok := d.GetOk("container.0.type"); ok {
		container := &marathon.Container{}
		t := v.(string)

		container.Type = t

		if t == "DOCKER" {
			docker := &marathon.Docker{}

			if v, ok := d.GetOk("container.0.docker.0.image"); ok {
				docker.Image = v.(string)
			}

			if v, ok := d.GetOk("container.0.docker.0.network"); ok {
				docker.Network = v.(string)
			}

			if v, ok := d.GetOk("container.0.docker.0.port_mappings.0.port_mapping.#"); ok {
				portMappings := make([]marathon.PortMapping, v.(int))

				for i, _ := range portMappings {
					portMappings[i] = marathon.PortMapping{}

					pmMap := d.Get(fmt.Sprintf("container.0.docker.0.port_mappings.0.port_mapping.%d", i)).(map[string]interface{})

					if val, ok := pmMap["container_port"]; ok {
						portMappings[i].ContainerPort = val.(int)
					}
					if val, ok := pmMap["host_port"]; ok {
						portMappings[i].HostPort = val.(int)
					}
					if val, ok := pmMap["protocol"]; ok {
						portMappings[i].Protocol = val.(string)
					}
					if val, ok := pmMap["service_port"]; ok {
						portMappings[i].ServicePort = val.(int)
					}

				}
				docker.PortMappings = portMappings
			}
			container.Docker = docker

		}

		if v, ok := d.GetOk("container.0.volumes.0.volume.#"); ok {
			volumes := make([]marathon.Volume, v.(int))

			for i, _ := range volumes {
				volumes[i] = marathon.Volume{}

				volumeMap := d.Get(fmt.Sprintf("container.0.volumes.0.volume.%d", i)).(map[string]interface{})

				if val, ok := volumeMap["container_path"]; ok {
					volumes[i].ContainerPath = val.(string)
				}
				if val, ok := volumeMap["host_path"]; ok {
					volumes[i].HostPath = val.(string)
				}
				if val, ok := volumeMap["mode"]; ok {
					volumes[i].Mode = val.(string)
				}
			}
			container.Volumes = volumes
		}

		appMutable.Container = container
	}

	if v, ok := d.GetOk("cpus"); ok {
		cpus, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.Cpus = cpus
	}

	if v, ok := d.GetOk("dependencies.#"); ok {
		dependencies := make([]string, v.(int))

		for i, _ := range dependencies {
			dependencies[i] = d.Get("dependencies." + strconv.Itoa(i)).(string)
		}

		if len(dependencies) != 0 {
			appMutable.Dependencies = dependencies
		}
	}

	if v, ok := d.GetOk("disk"); ok {
		disk, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.Disk = disk
	}

	if v, ok := d.GetOk("env"); ok {
		envMap := v.(map[string]interface{})
		env := make(map[string]string, len(envMap))

		for k, v := range envMap {
			env[k] = v.(string)
		}

		appMutable.Env = env
	}

	if v, ok := d.GetOk("health_checks.0.health_check.#"); ok {
		healthChecks := make([]marathon.HealthCheck, v.(int))

		for i, _ := range healthChecks {
			healthCheck := marathon.HealthCheck{}
			mapStruct := d.Get("health_checks.0.health_check." + strconv.Itoa(i)).(map[string]interface{})

			if prop, ok := d.GetOk("health_checks.0.health_check." + strconv.Itoa(i) + ".command.value"); ok {
				command := make(map[string]string)
				command["value"] = prop.(string)

				healthCheck.Command = command
			}

			if prop, ok := mapStruct["grace_period_seconds"]; ok {
				healthCheck.GracePeriodSeconds = prop.(int)
			}

			if prop, ok := mapStruct["interval_seconds"]; ok {
				healthCheck.IntervalSeconds = prop.(int)
			}

			if prop, ok := mapStruct["max_consecutive_failures"]; ok {
				healthCheck.MaxConsecutiveFailures = prop.(int)
			}

			if prop, ok := mapStruct["path"]; ok {
				healthCheck.Path = prop.(string)
			}

			if prop, ok := mapStruct["port_index"]; ok {
				healthCheck.PortIndex = prop.(int)
			}

			if prop, ok := mapStruct["protocol"]; ok {
				healthCheck.Protocol = prop.(string)
			}

			if prop, ok := mapStruct["timeout_seconds"]; ok {
				healthCheck.TimeoutSeconds = prop.(int)
			}

			healthChecks[i] = healthCheck
		}

		appMutable.HealthChecks = healthChecks
	}

	if v, ok := d.GetOk("instances"); ok {
		appMutable.Instances = v.(int)
	}

	if v, ok := d.GetOk("mem"); ok {
		mem, _ := strconv.ParseFloat(v.(string), 64)
		appMutable.Mem = mem
	}

	if v, ok := d.GetOk("ports.#"); ok {
		ports := make([]int, v.(int))

		for i, _ := range ports {
			ports[i] = d.Get("ports." + strconv.Itoa(i)).(int)
		}

		if len(ports) != 0 {
			appMutable.Ports = ports
		}
	}

	if v, ok := d.GetOk("upgrade_strategy.minimum_health_capacity"); ok {
		capacity, _ := strconv.ParseFloat(v.(string), 64)

		upgradeStrategy := &marathon.UpgradeStrategy{
			MinimumHealthCapacity: capacity,
		}
		appMutable.UpgradeStrategy = upgradeStrategy

	}

	if v, ok := d.GetOk("uris.#"); ok {
		uris := make([]string, v.(int))

		for i, _ := range uris {
			uris[i] = d.Get("uris." + strconv.Itoa(i)).(string)
		}

		if len(uris) != 0 {
			appMutable.Uris = uris
		}
	}

	return appMutable
}
