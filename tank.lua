require "cannon"
require "debugFuncs"
require "vec2"

function Tank(id, x, y)
  local t = {}
  t.id = id
  t.w = 30
  t.h = 50
  t.pos = Vec2(x+(t.w/2), y+(t.h/2))
  t.angle = 0
  t.angularVelocity = 0
  t.pointOfRotation = Vec2(0, 0)  -- Offset from middle
  t.velMag = 0 -- Magnitude of acceleration
  t.vel = Vec2(0, 0)
  t.cannon = Cannon(t)

  function t:applyResistance(dt)  -- Due to friction/air resistance
    self.velMag = self.velMag*(TANK_DECEL*math.min(144/dt, 1))
  end

  function t:accelerate(dt, totalPower)   -- Can be used to do both tracks rather than doing it individually.
    self.velMag = self.velMag + (TANK_ACC*dt)*totalPower
  end

  function t:leftTrack(dt, power)
    self.pointOfRotation = Vec2(self.w/2, 0)
    self.angularVelocity = (TANK_ROTATE_RATE*dt)*math.abs(power)
    if power > 0 then
      self.angle = self.angle - self.angularVelocity
    else
      self.angle = self.angle + self.angularVelocity -- If reversing
    end
    self:accelerate(dt, power)
    self.pointOfRotation = Vec2(0, 0)
  end

  function t:rightTrack(dt, power)
    self.pointOfRotation = Vec2(-self.w/2, 0)
    self.angularVelocity = (TANK_ROTATE_RATE*dt)*math.abs(power)
    if power > 0 then
      self.angle = self.angle + self.angularVelocity
    else
      self.angle = self.angle - self.angularVelocity -- If reversing
    end
    self:accelerate(dt, power)
    self.pointOfRotation = Vec2(0, 0)
  end

  function t:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.translate(self.pos.x+self.pointOfRotation.x, self.pos.y+self.pointOfRotation.y)
      lg.rotate(-self.angle)
      lg.translate(-self.pointOfRotation.x, -self.pointOfRotation.y)
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
      lg.line(0, 20, 0, self.h/2)  -- For knowing orientation.

      self.cannon:draw()  -- Will be drawn with current transformations.
    lg.pop()

    if VEL_DEBUG then
      debugVel(self)
    end
  end

  function t:update(dt)
    if love.keyboard.isDown("w", "up") then
      if love.keyboard.isDown("a", "left") then
        self:leftTrack(dt, 0.1)
        self:rightTrack(dt, 0.9)
      elseif love.keyboard.isDown("d", "right") then
        self:leftTrack(dt, 0.9)
        self:rightTrack(dt, 0.1)
      else
        self:accelerate(dt, 1)
      end
    elseif love.keyboard.isDown("s", "down") then
      if love.keyboard.isDown("a", "left") then
        self:leftTrack(dt, -1)
        self:rightTrack(dt, 1)
      elseif love.keyboard.isDown("d", "right") then
        self:leftTrack(dt, 1)
        self:rightTrack(dt, -1)
      else
        self:accelerate(dt, -1)
      end
    elseif love.keyboard.isDown("a", "left") and love.keyboard.isDown("d", "right") then
      self:accelerate(dt, 1)
    elseif love.keyboard.isDown("a", "left") then
      self:rightTrack(dt, 1)
    elseif love.keyboard.isDown("d", "right") then
      self:leftTrack(dt, 1)
    end

    self:applyResistance(dt)

    self.vel.x, self.vel.y = self.velMag*math.sin(self.angle), self.velMag*math.cos(self.angle)
    self.pos.x = self.pos.x + self.vel.x*dt
    self.pos.y = self.pos.y + self.vel.y*dt
    self.cannon:update(dt)
    self.angularVelocity = 0
  end

  table.insert(entities.players, t)
  return t
end
