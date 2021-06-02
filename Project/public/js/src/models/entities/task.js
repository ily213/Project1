
export class Task {
    constructor (id, project, type, priority, employeer, condition, name, desc, time, facttime){
        this.Id = id
        this.Project = project
        this.Type = type
        this.Priority = priority
        this.Employeer = employeer
        this.Condition = condition
        this.Name = name
        this.Desc = desc
        this.Time = time
        this.Facttime = facttime
    }
}

export let TASK_TYPE = {
    test: 'тест',
    create: 'создание',
    fix: 'исправление ошибок',
    change: 'изменение',
}

export let TASK_PRIORITY = {
    low: 'низкий',
    medium: 'средний',
    high: 'высокий',
    important: 'важно',

}

export let TASK_CONDITION = {
    new: 'новая',
    given: 'назначена',
    inwork: 'в работе',
    pause: 'пауза',
    complete: 'решено',
    tochange: 'согласовано',
}