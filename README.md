// Write about branching

Useful commands :

export $(grep -v '^#' .env | xargs) && aws dynamodb create-table --cli-input-json file://deployments/schema/users.json

export $(grep -v '^#' .env | xargs) && aws dynamodb list-tables --endpoint-url http://localhost:8001

db.users.find({ '$or' : [{ 'traits.id' : '1' }, { 'traits.id' : '8' }]  }).count()

db.users.aggregate([{ '$project' : { total_value : { '$sum' : '$traits.value' }  }   }]).sort({ 'total_value': -1  })



db.users.aggregate([{ 
'$match' : { 
  '$or': [
    {'traits.id' : '1'}, {'traits.id' : '4'} ] },
'$project' : { total_value : { '$sum' : '$traits.value' }}   
}]).sort({ 'total_value': -1  })



db.users.aggregate([{
  '$match' : {
    '$or': [
      {'traits.id' : '1'}, 
      {'traits.id' : '4'} 
    ] 
  },
}]).sort({ 'total_value': -1  })

-------
this one should work

// todo : add where id not = mine
db.users.aggregate([
  {
    '$match' : {
      '$and': [
        { 'gender': 'male' },
        { 'age' : { '$lte': 100 } },
        { 
          '$or' : [
              {'traits.id' : '1'},
              {'traits.id' : '4'}
            ] 
          }
      ]
    },
  },
  {
    '$project' :  { 
      attractiveness_score : { '$sum' : '$traits.value' },
      name: '$name',
      email: '$email',
      age: '$age',
      gender: '$gender',
      traits: '$traits'
    }
  },
  {
     $count: "email"
   }
]).sort({ 'attractiveness_score' : -1 })
------------
[//]: # (.sort&#40;{ 'total_value': -1  }&#41;)

